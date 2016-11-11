package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

type Entry struct {
	Name string
	Type string
	Size int
	Desc string
}

func (e *Entry) ToBuff() string {
	if e.Type == "string" {
		if e.Size < 0 {
			panic("ToBuff: string size error")
		}
		return `if len(log.` + e.Name + `) > ` + strconv.Itoa(e.Size) + ` {
		buff = append(buff, []byte(log.` + e.Name + `[:` + strconv.Itoa(e.Size) + `])...)
	} else {
		buff = append(buff, []byte(log.` + e.Name + `)...)
	}`
	} else if e.Type == "int32" || e.Type == "int64" {
		return "buff = strconv.AppendInt(buff, int64(log." + e.Name + "), 10)"
	} else if e.Type == "uint32" || e.Type == "uint64" {
		return "buff = strconv.AppendUint(buff, uint64(log." + e.Name + "), 10)"
	} else if e.Type == "time.Time" {
		return "buff = append(buff, []byte(log." + e.Name + ".Format(layout))...)"
	} else if e.Type == "float64" {
		return "buff = strconv.AppendFloat(buff, float64(log." + e.Name + "), 'g', 10, 64)"
	}
	panic(fmt.Sprintf("ToBuff: type not defined [%s]", e.Type))
	return ""
}

type NewStruct struct {
	Name     string
	Desc     string
	Entry    []*Entry
	BuffSize int
}

func (s *NewStruct) GetEntrys() []*Entry {
	var entrys []*Entry
	for _, e := range s.Entry {
		// 过滤最基本的接口请求字段，Packet()时自动赋值
		if e.Name == "GameSvrId" ||
			e.Name == "VGameAppid" ||
			e.Name == "PlatID" ||
			e.Name == "DtEventTime" {
			continue
		}
		entrys = append(entrys, e)
	}
	return entrys
}

func (s *NewStruct) SetDefaultParams() string {
	if s.Name == "GameSvrState" || s.Name == "IDIPFLOW" {
		return `
	s_` + s.Name + `.DtEventTime = time.Now()
	`
	} else {
		return `
	s_` + s.Name + `.GameSvrId = g_GameSrvID
	s_` + s.Name + `.VGameAppid = g_GameAppID
	s_` + s.Name + `.PlatID = g_GamePlatID
	s_` + s.Name + `.DtEventTime = time.Now()
	`
	}
}

func (s *NewStruct) ToParams() string {
	var params []byte
	for _, e := range s.GetEntrys() {
		params = append(params, []byte(e.Name+" "+e.Type+", ")...)
	}
	return string(params[:len(params)-2])
}

func (s *NewStruct) ToTestParams() string {
	var params []byte
	for _, e := range s.GetEntrys() {
		val := ""
		switch e.Type {
		case "string":
			val = `"test string"`
		case "uint32":
			val = "1234567"
		case "uint64":
			val = "1234567890"
		case "int32":
			val = "123"
		case "int64":
			val = "12345"
		case "time.Time":
			val = "time.Now()"
		case "float64":
			val = "123.456"
		}
		params = append(params, []byte(val+", ")...)
	}
	return string(params[:len(params)-2])
}

func (s *NewStruct) GetSize() int {
	size := len(s.Name) + 1
	for _, e := range s.Entry {
		if e.Type == "time.Time" {
			size += 19
		} else if e.Type == "string" {
			size += e.Size
		} else {
			size += 8
		}
		size += 1
	}
	return size
}

var (
	structs    []*NewStruct
	cur_struct *NewStruct
	tpl        = `// 本文件由 tlog-parser 根据 日志xml规范生成
// 请勿手动改写！

package tlog
import (
	"time"
	"strconv"
	"game_server/mdb"
	glog "core/log"
)

var layout = "2006-01-02 15:04:05"
{{range $s:=.}}
// 此结构体由工具按XML描述自动生成
type {{$s.Name}} struct {
{{range $e:=.Entry}}	{{$e.Name}} {{$e.Type}} // {{$e.Desc}}
{{end}}
}

// Log接口要求的Packet方法
func (log *{{$s.Name}}) Packet() []byte {
	struct_name := "{{$s.Name}}"
	buff := make([]byte, 0, {{$s.GetSize}})
	buff = append(buff, []byte(struct_name)...)
	{{range $e:=.Entry}}
	buff = append(buff, '|')
	{{$e.ToBuff}}
	{{end}}
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *{{$s.Name}}) InvokeHook() {}
func (log *{{$s.Name}}) CommitToTLog() {}
func (log *{{$s.Name}}) CommitToXdLog() {}
func (log *{{$s.Name}}) CommitToMySql() error { return nil }
func (log *{{$s.Name}}) Rollback() {}
func (log *{{$s.Name}}) Free() {}
func (log *{{$s.Name}}) SetEventId(time.Time) {}

func (log *{{$s.Name}}) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *{{$s.Name}}) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] {{$s.Name}} error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func New{{$s.Name}}({{$s.ToParams}}) *{{$s.Name}} {
	s_{{$s.Name}}:=new({{$s.Name}})
	{{$s.SetDefaultParams}}
	{{range $e:=.GetEntrys}}
	s_{{$s.Name}}.{{$e.Name}}={{$e.Name}}{{end}}
	return s_{{$s.Name}}
}
{{end}}
`
	tpl_test = `// 本文件由 tlog-parser 根据 日志xml规范生成
// 请勿手动改写！
package tlog

import (
	"testing"
	"os"
	"time"
)

var logs []Log

func Test(t *testing.T) {
	{{range $s:=.}}
	logs=append(logs,New{{$s.Name}}({{$s.ToTestParams}}))
	{{end}}
	f, err := os.Create("../../testdata/log.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, log := range logs {
		debugWriteToFile(log, f)
	}
}`
)

/*
{{range $s:=.}}
// 此结构体由工具按XML描述自动生成
type {{$s.Name}} struct {
{{range $e:=.Entry}}	{{$e.Name}} {{$e.Type}} // {{$e.Desc}}
{{end}}
*/

//读取日志xml规范
func read_xml(path string) error {
	f, err := os.Open(path)
	if err != nil {
		log.Println("read_xml:OpenFile:", err)
		return err
	}
	dc := xml.NewDecoder(f)
	for t, err := dc.Token(); err == nil; t, err = dc.Token() {
		if v, ok := t.(xml.StartElement); ok {
			if v.Name.Local == "struct" {
				cur_struct = new(NewStruct)
				structs = append(structs, cur_struct)
				for _, attr := range v.Attr {
					switch attr.Name.Local {
					case "name":
						cur_struct.Name = attr.Value
					case "desc":
						cur_struct.Desc = attr.Value
					}
				}
			} else if v.Name.Local == "entry" {
				entry := new(Entry)
				cur_struct.Entry = append(cur_struct.Entry, entry)
				for _, attr := range v.Attr {
					switch attr.Name.Local {
					case "name":
						tmp := strings.SplitAfterN(attr.Value, "", 2)
						tmp[0] = strings.ToUpper(tmp[0])
						entry.Name = strings.Join(tmp, "")
					case "type":
						switch attr.Value {
						case "string":
							entry.Type = "string"
						case "biguint":
							entry.Type = "uint64"
						case "uint":
							entry.Type = "uint32"
						case "int":
							entry.Type = "int32"
						case "float":
							entry.Type = "float64"
						case "datetime":
							entry.Type = "time.Time"
						default:
							return errors.New(fmt.Sprintf("read_xml: type [%s] not defined.", attr.Name.Local))
						}
					case "size":
						if len(attr.Value) > 0 {
							size, e := strconv.Atoi(attr.Value)
							if e != nil {
								return e
							}
							entry.Size = size
						} else {
							entry.Size = -1
						}
					case "desc":
						entry.Desc = attr.Value
					}
				}
			}
		}
	}
	//读取xml文件并检查标签数量
	return nil
}

//遍历xml文件夹
func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		log.Println("visit:", err)
		return err
	}
	if f.IsDir() {
		return nil
	}
	return read_xml(path)
}

//写入tlog-ext.go文件
func write(path string) {
	tmpl, err := template.New("tlog-ext").Parse(tpl)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(filepath.Join(path, "tlog-ext.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, structs)
	if err != nil {
		panic(err)
	}

	//test
	tmpl2, err := template.New("tlog-ext_test").Parse(tpl_test)
	if err != nil {
		panic(err)
	}
	f2, err := os.Create(filepath.Join(path, "tlog-ext_test.go"))
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	err = tmpl2.Execute(f2, structs)
	if err != nil {
		panic(err)
	}
}

//main
func main() {
	dir_path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(filepath.Join(dir_path, "../xml"), visit)
	if err != nil {
		log.Println("main:", err)
	}
	write(filepath.Join(dir_path, "../src/tlog"))
}
