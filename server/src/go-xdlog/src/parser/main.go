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
	Name     string
	Type     string
	Size     int
	Desc     string
	JsonName string
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
	TagName  string
	Desc     string
	Entry    []*Entry
	BuffSize int
}

func (s *NewStruct) GetEntrys() []*Entry {
	var entrys []*Entry
	for _, e := range s.Entry {
		// 过滤最基本的接口请求字段，Packet()时自动赋值
		if e.Name == "Sid" ||
			e.Name == "Time" {
			continue
		}
		entrys = append(entrys, e)
	}
	return entrys
}

func (s *NewStruct) SetDefaultParams() string {
	if s.Name == "SnapshotFlow" {
		return `
		s_` + s.Name + `.Sid = g_GameSrvID`
	} else {
		return `
		s_` + s.Name + `.Sid = g_GameSrvID
		s_` + s.Name + `.Time = gotime.GetNowTime()`

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
	tpl        = `// 本文件由 xdlog-parser 根据日志xml规范生成
//请勿手动改写！
package xdlog

import (
	"game_server/mdb"
	"encoding/json"
	gotime "core/time"
	//"strconv"
	"fmt"
	//"core/fail"
	"time"
	//."game_server/config"
)

var layout = "2006-01-02 15:04:05"

{{range $s:=.}}
// 此结构体由工具按XML描述自动生成
type {{$s.Name}} struct {
	Tag string Json好麻烦json:"tag"Json好麻烦
	Eid uint64 Json好麻烦json:"eid"Json好麻烦
	Microtime string Json好麻烦json:"microtime"Json好麻烦
{{range $e:=.Entry}}	{{$e.Name}}	{{$e.Type}} Json好麻烦json:"{{$e.JsonName}}"Json好麻烦 // {{$e.Desc}}
{{end}}
}

// Log接口要求的Packet方法
func (log *{{$s.Name}}) Packet() []byte {
	log.Tag = "{{$s.TagName}}"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *{{$s.Name}}) InvokeHook() {}
func (log *{{$s.Name}}) CommitToTLog() {}
func (log *{{$s.Name}}) CommitToXdLog() {}
func (log *{{$s.Name}}) CommitToMySql() error { return nil }
func (log *{{$s.Name}}) Rollback() {}
func (log *{{$s.Name}}) Free() {}

func (log *{{$s.Name}}) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *{{$s.Name}}) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *{{$s.Name}}) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
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

/*
func createeid(time time.Time) uint64 {
	timenanoid := time.UnixNano()
	if ServerCfg.EnableGlobalServer {
		eid, err := strconv.ParseUint(fmt.Sprintf("%v", timenanoid)+"1", 10, 64)
		fail.When(err != nil, err)
		return eid
	}
	eid, err := strconv.ParseUint(fmt.Sprintf("%v", timenanoid)+"0", 10, 64)
	fail.When(err != nil, err)
	return eid
}
*/
`

	tpl_test = `// 本文件由 xdlog-parser 根据日志xml规范生成
//请勿手动改写！
package xdlog

import (
	"time"
	"strconv"
	"game_server/mdb"
	glog "core/log"
	"encoding/json"
	gotime "core/time"
	"strings"
)

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

func first2Upper(str string) string {
	if len(str) == 0 {
		return str
	}
	first := strings.SplitAfterN(str, "", 2)[0]
	other := str[1:]
	first = strings.ToUpper(first)
	return first + other
}

func first2Lower(str string) string {
	if len(str) == 0 {
		return str
	}
	first := strings.SplitAfterN(str, "", 2)[0]
	other := str[1:]
	first = strings.ToLower(first)
	return first + other
}

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
						cur_struct.TagName = strings.Replace(first2Lower(cur_struct.Name), "Flow", "", -1)
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
						entry.Name = first2Upper(attr.Value)
						entry.JsonName = strings.ToLower(attr.Value)
					case "type":
						switch attr.Value {
						case "string":
							entry.Type = "string"
						case "biguint":
							entry.Type = "uint64"
						case "bigint":
							entry.Type = "int64"
						case "uint":
							entry.Type = "uint32"
						case "int":
							entry.Type = "int32"
						case "float":
							entry.Type = "float64"
						case "datetime":
							entry.Type = "time.Time"
						default:
							return errors.New(fmt.Sprintf("read_xml: type [%s] not found", attr.Name.Local))
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

//写入xdlog-ext.go文件
func write(path string) {
	tmpl, err := template.New("xdlog-ext").Parse(strings.Replace(tpl, "Json好麻烦", "`", -1))
	if err != nil {
		panic(err)
	}
	f, err := os.Create(filepath.Join(path, "xdlog-ext.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, structs)
	if err != nil {
		panic(err)
	}

	//test
	tmpl2, err := template.New("xdlog-ext_test").Parse(tpl_test)
	if err != nil {
		panic(err)
	}
	f2, err := os.Create(filepath.Join(path, "xdlog-xxd_test.go"))
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
	write(filepath.Join(dir_path, "../src/xdlog"))
}
