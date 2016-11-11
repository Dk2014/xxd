package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

// 指定目录 发现中文
func main() {
	var path string
	flag.StringVar(&path, "path", "", "the path of code")
	flag.Parse()
	if path != "" {
		_, err := os.Stat(path)
		fail(os.IsNotExist(err), errors.New("path not exists"))
		fail(err != nil, err)
		err = pathWalk(path)
		fail(err != nil, err)
	}
}

func fail(condition bool, err error) {
	if condition {
		panic(err)
	}
}

func pathWalk(path string) error {
	// 遍历所有文件
	err := filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		regFile, _ := regexp.Compile(".+.go$")
		if fi.IsDir() {
			return nil
		} else {
			if regFile.Match([]byte(path)) {
				err := handle(path)
				return err
			}
			return nil
		}
	})
	return err
}

func handle(path string) error {
	// 处理文件 找出中文
	pattern := "([\u4e00-\u9fa5]+)"
	//note_pattern := "//\s+"
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	bfrd := bufio.NewReader(f)
	line_num := 1
	for {
		line, err := bfrd.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		results := reg.FindAll(line, -1)
		if len(results) > 0 {
			fmt.Printf("%v at line %d\n", path, line_num)
			fmt.Println(string(line))
		}
		line_num++
	}
	return nil
}
