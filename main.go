package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			path := os.Args[i]
			TransPy(path)
		}
	} else {
		ScanDir("./res/")
		ScanDir("./")
	}
}

func ScanDir(dir string) {
	files, err := ListDir(dir, ".py")
	if err != nil {
		//fmt.Println("ListDir error:", err)
		return
	}
	for _, path := range files {
		TransPy(path)
	}
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

func TransPy(path string) error {
	fmt.Println("translate:", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile fail:", path, err)
		return err
	}

	buf := bytes.NewBuffer(data)

	var part = new(CodePart)
	err = part.Parse(buf)
	if err != nil {
		fmt.Println("Parse fail:", err)
		return err
	}

	r := part.Translate()

	last := strings.LastIndex(path, ".py")
	newPath := path[0:last] + ".go"
	fmt.Println("write file:", newPath)
	err = ioutil.WriteFile(newPath, []byte(r), os.ModeType)
	if err != nil {
		fmt.Println("WriteFile error:", err)
	}
	return err
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}
