package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
	文件格式
	文件头：							|uint32|
	表名+字段名：第一条[]byte:			|tableName,field1,field2,field3,...|
	字段类型：第二条[]byte:			|uint8|uint8|uint8|...|		(1-char, 2-varchar, 3-int, 4-bigint)
	协议数据： 第三条[]byte: 			|模块ID, 模块名, actionID, action名, action...|

	协议号：							|uint32|  低16为是网络协议uint8|uint8,分别是模块id和接口id；高16位自定义常量标识
	事务发生时间：					|int64|

	事务头：							|uint32|
	操作标识：						|uint8| 					(1-插入，2-删除，3-修改)
	表序号：							|uint32|					(序号对应init()中的表信息记录序列，记录事务处理所属哪个表)
	内容：							|bytes|
*/

var (
	ComBytes = []byte{','}
	IdBytes  = []byte("id")
	PidBytes = []byte("pid")
)

type TypeCode int

const (
	MYSQL_TINYINT    = TypeCode(1)
	MYSQL_SMALLINT   = TypeCode(2)
	MYSQL_INT        = TypeCode(3)
	MYSQL_BIGINT     = TypeCode(4)
	MYSQL_TEXT       = TypeCode(5)
	MYSQL_CHAR       = TypeCode(6)
	MYSQL_VARCHAR    = TypeCode(7)
	MYSQL_BLOB       = TypeCode(8)
	MYSQL_BINARY     = TypeCode(9)
	MYSQL_FLOAT      = TypeCode(10)
	MYSQL_DOUBLE     = TypeCode(11)
	MYSQL_TIMESTAMP  = TypeCode(12)
	MYSQL_MEDIUMBLOB = TypeCode(13)
)

const (
	OP_INSERT = 1
	OP_DELETE = 2
	OP_UPDATE = 3
)

func main() {
	var (
		filename           = flag.String("file", "", "The sync file path.")
		plrid              = flag.String("plrid", "", "The player ID the result filtering by.")
		tableFilterPtr     = flag.String("tables", "", "The specifical tables will be traced on. (Separated by colon)")
		tableFilter        string
		timeBeginStr       = flag.String("begin", "", "Time begin.")
		timeEndStr         = flag.String("end", "", "Time end.")
		timeBegin, timeEnd int64
	)

	flag.Parse()

	if *tableFilterPtr != "" {
		tableFilter = "," + *tableFilterPtr + ","
	} else {
		tableFilter = ""
	}
	if *timeBeginStr != "" {
		t, err := time.Parse("2006-01-02 03:04:05 -0700 MST", *timeBeginStr)
		if err != nil {
			timeBegin = 0
		} else {
			timeBegin = t.Unix()
		}
	} else {
		timeBegin = 0
	}
	if *timeEndStr != "" {
		t, err := time.Parse("2006-01-02 03:04:05 -0700 MST", *timeEndStr)
		if err != nil {
			timeEnd = 0
		} else {
			timeEnd = t.Unix()
		}
	} else {
		timeEnd = 0
	}
	reader, err := NewReader(*filename)
	if err != nil {
		println("Open sync file failed:", err)
		return
	}
	defer reader.Close()

	// table and column info
	head1, err1 := reader.ReadChunk()
	if err1 != nil {
		println("Read sync file head1 failed:", err1)
		return
	}

	// protocol info
	head2, err2 := reader.ReadChunk()
	if err2 != nil {
		println("Read sync file head2 failed:", err2)
		return
	}

	// rpc info
	head3, err3 := reader.ReadChunk()
	if err3 != nil {
		println("Read sync file head3 failed:", err3)
		return
	}

	syncInfo := NewSyncInfo(head1, head2, head3)

	for {
		chunk, err := reader.ReadChunk()
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Read chunk failed: %v\n", err)
			}
			break
		}

		protoId := chunk.ReadUInt32()
		timestamp := chunk.ReadInt64()

		printedStamp := false

		for chunk.Length() > 0 {
			operator := chunk.ReadUInt8()
			tableId := chunk.ReadInt32()

			table := syncInfo.GetTable(tableId)

			switch operator {
			case OP_INSERT:
				values := table.Values(chunk)
				if (*plrid == "" || values[table.PlrIDName()] == *plrid) && (tableFilter == "" || strings.Contains(tableFilter, ","+table.Name+",")) && (timeBegin == 0 || timestamp >= timeBegin) && (timeEnd == 0 || timestamp < timeEnd) {
					if !printedStamp {
						fmt.Printf("// %d - %s", protoId, time.Unix(timestamp, 0).String())
						fmt.Println()
						printedStamp = true
					}

					fmt.Print("REPLACE INTO `", table.Name, "` (", table.ColumnList, ") VALUES (")
					values.Print(&table)
					fmt.Print(");")
					fmt.Println()

					fmt.Print("// ")
					fmt.Print("DELETE FROM `", table.Name, "` WHERE `", table.KeyName, "` = ", values[table.KeyName], ";")
					fmt.Println()
					fmt.Println("// Insert")
				}
			case OP_DELETE:
				values := table.Values(chunk)
				if (*plrid == "" || values[table.PlrIDName()] == *plrid) && (tableFilter == "" || strings.Contains(tableFilter, ","+table.Name+",")) && (timeBegin == 0 || timestamp >= timeBegin) && (timeEnd == 0 || timestamp < timeEnd) {
					if !printedStamp {
						fmt.Printf("// %d - %s", protoId, time.Unix(timestamp, 0).String())
						fmt.Println()
						printedStamp = true
					}

					fmt.Print("// ")
					fmt.Print("REPLACE INTO `", table.Name, "` (", table.ColumnList, ") VALUES (")
					values.Print(&table)
					fmt.Print(");")
					fmt.Println()

					fmt.Print("DELETE FROM `", table.Name, "` WHERE `", table.KeyName, "` = ", values[table.KeyName], ";")
					fmt.Println()
					fmt.Println("// Delete")
				}
			case OP_UPDATE:
				valuesNew := table.Values(chunk)
				valuesOld := table.Values(chunk)
				if (*plrid == "" || valuesNew[table.PlrIDName()] == *plrid || valuesOld[table.PlrIDName()] == *plrid) && (tableFilter == "" || strings.Contains(tableFilter, ","+table.Name+",")) && (timeBegin == 0 || timestamp >= timeBegin) && (timeEnd == 0 || timestamp < timeEnd) {
					if !printedStamp {
						fmt.Printf("// %d - %s", protoId, time.Unix(timestamp, 0).String())
						fmt.Println()
						printedStamp = true
					}

					fmt.Print("REPLACE INTO `", table.Name, "` (", table.ColumnList, ") VALUES (")
					valuesNew.Print(&table)
					fmt.Print(");")
					fmt.Println()

					fmt.Print("// ")
					fmt.Print("REPLACE INTO `", table.Name, "` (", table.ColumnList, ") VALUES (")
					valuesOld.Print(&table)
					fmt.Print(");")
					fmt.Println()
					for k, vNew := range valuesNew {
						if vOld := valuesOld[k]; vOld != vNew {
							fmt.Printf("// Update %v : %v -> %v", k, vOld, vNew)
							fmt.Println()
						}
					}
				}
			}
		}

		if printedStamp {
			fmt.Println()
		}
	}
}

type SyncInfo struct {
	Tables []TableInfo
}

func NewSyncInfo(head1, head2, head3 *Chunk) *SyncInfo {
	syncInfo := &SyncInfo{
		Tables: make([]TableInfo, 0, 50),
	}
	// parse table info
	id := int32(0)
	for {
		nameBytes := head1.ReadBytes()
		typeBytes := head1.ReadBytes()

		names := bytes.Split(nameBytes, ComBytes)

		columnBytes := nameBytes[bytes.Index(nameBytes, ComBytes)+1:]

		table := TableInfo{
			Id:         id,
			Name:       string(names[0]),
			ColumnList: string(columnBytes),
			Columns:    make([]ColumnInfo, len(names)-1),
		}

		for i := 0; i < len(table.Columns); i++ {
			if bytes.Equal(names[i+1], IdBytes) {
				table.KeyName = "id"
			} else if bytes.Equal(names[i+1], PidBytes) && table.KeyName != "id" {
				table.KeyName = "pid"
			}
			table.Columns[i] = ColumnInfo{Name: string(names[i+1]), Type: TypeCode(typeBytes[i])}
		}

		syncInfo.Tables = append(syncInfo.Tables, table)

		if head1.Length() == 0 {
			break
		}

		id += 1
	}

	// parse protocol info

	return syncInfo
}

func (si *SyncInfo) GetTable(id int32) TableInfo {
	for _, table := range si.Tables {
		if table.Id == id {
			return table
		}
	}
	return TableInfo{}
}

type TableInfo struct {
	Id         int32
	Name       string
	KeyName    string
	Columns    []ColumnInfo
	ColumnList string
}

type TableValues map[string]string

func (tInfo *TableInfo) PlrIDName() string {
	if tInfo.Name == "player" {
		return "id"
	} else {
		return "pid"
	}
}

func (tInfo *TableInfo) Values(chunk *Chunk) TableValues {
	tv := make(TableValues)
	for _, column := range tInfo.Columns {
		tv[column.Name] = chunk.ReadValue(column.Type)
	}
	return tv
}

func (tv TableValues) Print(tInfo *TableInfo) {
	for i, column := range tInfo.Columns {
		fmt.Print(tv[column.Name])
		if i < len(tInfo.Columns)-1 {
			fmt.Print(",")
		}
	}
}

func (ti *TableInfo) Values0(chunk *Chunk) string {
	var primaryKey string
	for i := 0; i < len(ti.Columns); i++ {
		value := chunk.ReadValue(ti.Columns[i].Type)
		fmt.Print(value)
		if ti.Columns[i].Name == ti.KeyName {
			primaryKey = value
		}
		if i < len(ti.Columns)-1 {
			fmt.Print(",")
		}
	}
	return primaryKey
}

type ColumnInfo struct {
	Name string
	Type TypeCode
}

type Reader struct {
	file *os.File
	head []byte
}

func NewReader(filename string) (*Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &Reader{
		file: file,
		head: make([]byte, 4),
	}, nil
}

func (reader *Reader) Close() {
	reader.file.Close()
}

func (reader *Reader) ReadChunk() (*Chunk, error) {
	if n, err := reader.file.Read(reader.head); n != len(reader.head) || err != nil {
		if err == nil && n != len(reader.head) {
			return nil, errors.New("incomplete chunk head")
		}
		return nil, err
	}
	data := make([]byte, binary.BigEndian.Uint32(reader.head))
	if n, err := reader.file.Read(data); n != len(data) || err != nil {
		if err == nil && n != len(data) {
			return nil, errors.New("incomplete chunk data")
		}
		return nil, err
	}
	return &Chunk{data: data}, nil
}

type Chunk struct {
	data []byte
}

func (chunk *Chunk) Length() int {
	return len(chunk.data)
}

func (chunk *Chunk) Read(n int) []byte {
	r := chunk.data[0:n]
	chunk.data = chunk.data[n:]
	return r
}

func (chunk *Chunk) ReadUInt8() uint8 {
	return uint8(chunk.Read(1)[0])
}

func (chunk *Chunk) ReadInt8() int8 {
	return int8(chunk.Read(1)[0])
}

func (chunk *Chunk) ReadByte() byte {
	return chunk.Read(1)[0]
}

func (chunk *Chunk) ReadUInt16() uint16 {
	return binary.BigEndian.Uint16(chunk.Read(2))
}

func (chunk *Chunk) ReadInt16() int16 {
	return int16(binary.BigEndian.Uint16(chunk.Read(2)))
}

func (chunk *Chunk) ReadUInt32() uint32 {
	return binary.BigEndian.Uint32(chunk.Read(4))
}

func (chunk *Chunk) ReadInt32() int32 {
	return int32(binary.BigEndian.Uint32(chunk.Read(4)))
}

func (chunk *Chunk) ReadUInt64() uint64 {
	return binary.BigEndian.Uint64(chunk.Read(8))
}

func (chunk *Chunk) ReadInt64() int64 {
	return int64(binary.BigEndian.Uint64(chunk.Read(8)))
}

func (chunk *Chunk) ReadFloat32() float32 {
	return math.Float32frombits(chunk.ReadUInt32())
}

func (chunk *Chunk) ReadFloat64() float64 {
	return math.Float64frombits(chunk.ReadUInt64())
}

func (chunk *Chunk) ReadBytes() []byte {
	return chunk.Read(int(binary.BigEndian.Uint32(chunk.Read(4))))
}

func (chunk *Chunk) ReadString() string {
	return string(chunk.ReadBytes())
}

func (chunk *Chunk) ReadValue(typeCode TypeCode) string {
	switch typeCode {
	case MYSQL_TINYINT:
		return strconv.Itoa(int(chunk.ReadInt8()))
	case MYSQL_SMALLINT:
		return strconv.Itoa(int(chunk.ReadInt16()))
	case MYSQL_INT:
		return strconv.Itoa(int(chunk.ReadInt32()))
	case MYSQL_BIGINT:
		return strconv.Itoa(int(chunk.ReadInt64()))
	case MYSQL_FLOAT:
		return strconv.FormatFloat(float64(chunk.ReadFloat32()), 'f', -1, 32)
	case MYSQL_DOUBLE:
		return strconv.FormatFloat(chunk.ReadFloat64(), 'f', -1, 64)
	case MYSQL_TIMESTAMP:
		return strconv.Itoa(int(chunk.ReadInt64()))
	case MYSQL_TEXT:
		fallthrough
	case MYSQL_CHAR:
		fallthrough
	case MYSQL_VARCHAR:
		fallthrough
	case MYSQL_BLOB:
		fallthrough
	case MYSQL_MEDIUMBLOB:
		fallthrough
	case MYSQL_BINARY:
		return "HEX('" + hex.EncodeToString(chunk.ReadBytes()) + "')"
	}
	panic(fmt.Sprintf("unknow type %#v", typeCode))
}
