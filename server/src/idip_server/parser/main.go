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

type NewStruct struct {
	Name     string
	Desc     string
	Entry    []*Entry
	Id       string
	IsHeader bool
}
type Macro struct {
	Name string
	Val  string
	Desc string
}
type Macrosgroup struct {
	Name  string
	Desc  string
	Macro []*Macro
}
type Process struct {
	Name string
	Req  string
	Rsp  string
}

var (
	structs         []*NewStruct
	cur_struct      *NewStruct
	macrosgroups    = make(map[string]*Macrosgroup)
	cur_macrosgroup *Macrosgroup
)

func getLen(name string) (int, error) {
	mg, exist := macrosgroups["NET_MACRO"]
	if !exist {
		panic("Can't found NET_MACRO")
	}
	for _, m := range mg.Macro {
		if m.Name == name {
			return strconv.Atoi(m.Val)
		}
	}
	return -1, errors.New("getLen:not found")
}

func getType(name string) error {
	for _, s := range structs {
		if s.Name == name {
			return nil
		}
	}
	return errors.New("getType:not found")
}

//读取xml
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
				flag := true
				for _, attr := range v.Attr {
					switch attr.Name.Local {
					case "name":
						if attr.Value == "IdipDataPaket" {
							flag = false
							break
						} else if attr.Value == "IdipHeader" {
							cur_struct.Name = "IDIP_REQ_HEAD"
							cur_struct.IsHeader = true
						} else {
							cur_struct.Name = attr.Value
							cur_struct.IsHeader = false
						}
					case "id":
						cur_struct.Id = attr.Value
					case "desc":
						cur_struct.Desc = attr.Value
					}
				}
				if flag {
					structs = append(structs, cur_struct)
				} else {
					cur_struct = nil
				}
			} else if v.Name.Local == "entry" {
				if cur_struct == nil {
					continue
				}
				entry := new(Entry)
				cur_struct.Entry = append(cur_struct.Entry, entry)
				for _, attr := range v.Attr {
					switch attr.Name.Local {
					case "name":
						entry.Name = attr.Value
					case "type":
						switch attr.Value {
						case "string":
							entry.Type = "string"
						case "uint64":
							entry.Type = "uint64"
						case "uint32":
							entry.Type = "uint32"
						case "uint8":
							entry.Type = "uint8"
						case "int32":
							entry.Type = "int32"
						default:
							if e := getType(attr.Value); e == nil {
								entry.Type = "[]" + attr.Value
							} else {
								return errors.New(fmt.Sprintf("read_xml: type [%s] not defined.", attr.Value))
							}
						}
					case "size":
						size, e := getLen(attr.Value)
						if e != nil {
							return e
						}
						entry.Size = size
					case "desc":
						entry.Desc = attr.Value
					}
				}
			} else if v.Name.Local == "macrosgroup" {
				cur_macrosgroup = new(Macrosgroup)
				for _, attr := range v.Attr {
					switch attr.Name.Local {
					case "name":
						cur_macrosgroup.Name = attr.Value
						macrosgroups[attr.Value] = cur_macrosgroup
					case "desc":
						cur_macrosgroup.Desc = attr.Value
					}
				}
			} else if v.Name.Local == "macro" {
				macro := new(Macro)
				flag := true
				for _, attr := range v.Attr {
					switch attr.Name.Local {
					case "name":
						if strings.HasSuffix(attr.Value, "_RSP") {
							flag = false
							break
						}
						macro.Name = attr.Value
					case "value":
						macro.Val = attr.Value
					case "desc":
						macro.Desc = attr.Value
					}
				}
				if flag {
					cur_macrosgroup.Macro = append(cur_macrosgroup.Macro, macro)
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

//写入.go文件
func write(path string) {
	//protocol.go
	t_protocol, err := template.ParseFiles(filepath.Join(path, "parser/protocol.tpl"))
	if err != nil {
		panic(err)
	}
	f_protocol, err := os.Create(filepath.Join(path, "protocol.go"))
	if err != nil {
		panic(err)
	}
	defer f_protocol.Close()
	err = t_protocol.Execute(f_protocol, structs)
	if err != nil {
		panic(err)
	}
	//routine.go
	t_routine, err := template.ParseFiles(filepath.Join(path, "parser/routine.tpl"))
	if err != nil {
		panic(err)
	}
	f_routine, err := os.Create(filepath.Join(path, "routine.go"))
	if err != nil {
		panic(err)
	}
	defer f_routine.Close()
	err = t_routine.Execute(f_routine, macrosgroups)
	if err != nil {
		panic(err)
	}
	//proc_*.go
	t_process, err := template.ParseFiles(filepath.Join(path, "parser/process.tpl"))
	if err != nil {
		panic(err)
	}
	mg, exist := macrosgroups["NET_CMD_ID"]
	if !exist {
		panic("Can't found NET_CMD_ID")
	}
	for _, m := range mg.Macro {
		if strings.HasSuffix(m.Name, "_REQ") {
			p := new(Process)
			p.Req = m.Name
			p.Name = m.Name[:len(m.Name)-4]
			p.Rsp = p.Name + "_RSP"
			f_process, err := os.OpenFile(filepath.Join(path, "proc_"+p.Name+".go"), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
			if err != nil {
				log.Println("Create file err:", p.Name, err)
				continue
			}
			err = t_process.Execute(f_process, p)
			if err != nil {
				panic(err)
			}
			f_process.Close()
		}
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
	write(filepath.Join(dir_path, "../src/idip_server"))
}
