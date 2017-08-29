package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

//key word
const (
	CLASS  = "class"
	FUNC   = "def"
	FOR    = "for"
	IF     = "if"
	ELSEIF = "elif"
	ELSE   = "else"
)

type CodePartType int

const (
	CodePart_class CodePartType = iota
	CodePart_func
	CodePart_for
	CodePart_if
	CodePart_elseif
	CodePart_else

	CodePart_raw

	CodePart_NONE = -1
	CodePart_Root = -10
)

var codePartHeadMap = map[CodePartType]string{
	CodePart_class:  CLASS,
	CodePart_func:   FUNC,
	CodePart_for:    FOR,
	CodePart_if:     IF,
	CodePart_elseif: ELSEIF,
	CodePart_else:   ELSE,
}

type TransFunc func(part *CodePart) string

//反向map
var codePartHeadMapRev = make(map[string]CodePartType)

func ReplaceKeys(str string) string {
	l := len(str)
	for k, v := range keywordMap {
		str = strings.Replace(str, k, v, l)
	}
	return str
}

type CodePart struct {
	partType CodePartType
	children []*CodePart
	parent   *CodePart
	//head     string
	data    string
	headPos int
	trans   ITranslate
}

func init() {
	for key, val := range codePartHeadMap {
		codePartHeadMapRev[val] = key
	}
}

func getHeadPos(line string) int {
	trimline := strings.TrimSpace(line)
	pos := strings.Index(line, trimline)
	return pos
}

func getHeadType(line string) (CodePartType, int, string) {
	for key, val := range codePartHeadMap {
		trimline := strings.TrimSpace(line)
		vals := val

		if strings.HasPrefix(trimline, vals) {
			pos := strings.Index(line, trimline)
			return key, pos, vals
		}
	}
	return CodePart_NONE, 0, ""
}

func (p *CodePart) Parse(buf *bytes.Buffer) error {
	buf.WriteString("\n\n")
	return p.doParse(buf, CodePart_Root, -1, "", nil)
	/*
		head, err := buf.ReadString('\n')
		if err == io.EOF {
			return nil
		}

		if err != nil {
			fmt.Errorf("%v", err)
		}

		if tp, headpos, vals := getHeadType(head); tp == CodePart_NONE {
			return errors.New(fmt.Sprintf("getHeadType fail", head))
		} else {
			return p.doParse(buf, tp, headpos, parseHead(head, vals), nil)
		}
	*/
}

func parseHead(head string, key string) string {
	//fmt.Println("parseHead pre:", head, []byte(head), len(head))
	head = strings.TrimSpace(head)
	head = strings.TrimPrefix(head, key)
	end := strings.LastIndex(head, ":")
	head = head[0:end]
	//fmt.Println("parseHead:", head, "@@@@", key, "@@@", end)
	return head

}

func (p *CodePart) doParse(buf *bytes.Buffer, setpart CodePartType, headpos int, head string, parent *CodePart) error {
	p.partType = setpart
	p.headPos = headpos
	p.data = head
	p.trans = GetGolangTrans(setpart) //to golang
	p.parent = parent
	if p.partType == CodePart_raw {
		return nil
	}
	for {
		//fmt.Println("buf:", len(buf.Bytes()))
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			//fmt.Println("to end:", line)
			break
		}
		//判断空行
		if len(strings.TrimSpace(line)) <= 1 {
			continue
		}
		newpos := getHeadPos(line)
		//判断是否子块内
		if newpos > p.headPos {
			//块内，自己解
			if tp, headpos, vals := getHeadType(line); tp != CodePart_NONE {
				//新块类型
				var newpart = CodePart{}
				newpart.doParse(buf, tp, headpos, parseHead(line, vals), p)
				p.children = append(p.children, &newpart)
				continue
			} else {
				//基础代码
				if len(line) > 1 {
					line = line[0 : len(line)-2]
				}
				var newpart = CodePart{}
				newpart.doParse(buf, CodePart_raw, newpos, strings.TrimSpace(line), p)
				p.children = append(p.children, &newpart)
				//var newpart = CodePart{partType: CodePart_raw, data: line,}
				//p.children = append(p.children, &newpart)
				//fmt.Println("parse line:", line)
			}
		} else {
			//fmt.Println("parse out:", line)
			//塞回去，让上层解
			var left []byte
			left = append(left, buf.Bytes()...)
			buf.Reset()
			buf.WriteString(line)
			buf.Write(left)
			return nil
		}

	}

	return nil
}

//缩进
func (p *CodePart) writeScale(buf *bytes.Buffer) {
	var inClass bool = false
	var pr *CodePart = p
	//fmt.Println("writeScale ", pr)
	for {
		pr = pr.parent
		if pr == nil {
			break
		}
		if pr.partType == CodePart_class {
			inClass = true
			break
		}
	}
	//fmt.Println("writeScale end", inClass)
	var lastPos = p.headPos
	if inClass {
		lastPos -= 4
	}
	for index := 0; index < lastPos; index++ {
		buf.WriteString(" ")
	}
}
func (p *CodePart) doTranslate(buf *bytes.Buffer) {
	//fmt.Println("doTranslate:", p, "  p:", p, "  trans:", p.trans)

	p.trans.WriteHead(p, buf)
	p.trans.WriteBodyPre(p, buf)
	for _, v := range p.children {
		v.doTranslate(buf)
	}

	p.trans.WriteBodyEnd(p, buf)
}

func (p *CodePart) Translate() string {
	var buf = bytes.NewBufferString("")
	p.doTranslate(buf)
	str := string(buf.Bytes())
	str = ReplaceKeys(str)
	return str
}

func (p *CodePart) Print() {
	printTree(p, 0)
}

func printTree(part *CodePart, blk int) {
	for i := 0; i < blk; i++ {
		fmt.Print("    ") //缩进
	}
	fmt.Printf("|—<%v:%s-%d>\n", part.partType, part.data, part.headPos) //打印"|—<id>"形式
	for i := 0; i < len(part.children); i++ {
		printTree(part.children[i], blk+1) //打印子树，累加缩进次数
	}
}
