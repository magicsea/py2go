package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPy2go(t *testing.T) {

	fmt.Print(':', '\r', '\n', "--")
	path := "res/test.py"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile fail:", path, err)
		return
	}
	//log.Info("read len=%d", len(data))
	//log.Info("%s", data)
	buf := bytes.NewBuffer(data)
	fmt.Println(buf.Len())

	var part = new(CodePart)
	err = part.Parse(buf)
	if err != nil {
		fmt.Println("Parse fail:", err)
		return
	}

	fmt.Println("tree:")
	part.Print()
	r := part.Translate()
	fmt.Println("result:")
	fmt.Println(r)

}

func TestFile(t *testing.T) {
	path := "res/test.py"
	err := TransPy(path)
	fmt.Println(err)
}

func TestDir(t *testing.T) {
	ScanDir("./res/")
	ScanDir("./")
}
