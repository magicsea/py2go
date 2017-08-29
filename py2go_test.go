package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPy2go(t *testing.T) {

	fmt.Print(':', '\r', '\n', "--")
	path := "res/utility.py"
	//path := "res/test.py"
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
	part.partType = CodePart_Root
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
	if err := ScanDir("./res/"); err != nil {
		fmt.Println("ScanDir error:", err)
		return
	}
	if err := ScanDir("./"); err != nil {
		fmt.Println("ScanDir error:", err)
		return
	}
}
func TestSimple(t *testing.T) {
	var list = []int{0, 1, 2, 3, 4}
	var n = 3
	list = append(list[0:n], list[n+1:]...)
	fmt.Println(list)

}
