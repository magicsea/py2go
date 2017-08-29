package main

import (
	"bytes"
	"fmt"
	"strings"
)

//替换关键字
var keywordMap = map[string]string{
	" True":       " true",
	" False":      " false",
	" not ":       " !",
	" and ":       "&&",
	" or ":        "||",
	"is not None": "!=nil",
	" is not ":    "!=",
	" is ":        "==",
	" None ":      "nil",
	"#":           "//",
}

type GoRoot struct {
}

func (this *GoRoot) WriteHead(part *CodePart, buf *bytes.Buffer) {
}
func (this *GoRoot) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {

}
func (this *GoRoot) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {

}

type GoClass struct {
}

func (this *GoClass) WriteHead(part *CodePart, buf *bytes.Buffer) {
	last := strings.LastIndex(part.data, "(")
	str := part.data[0:last]
	//fmt.Println("str:%d %d %s", last, len(part.data), str)
	temp := "\ntype ##head## struct {}\n"
	temp = strings.Replace(temp, "##head##", str, 255)
	buf.WriteString(temp)
}
func (this *GoClass) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {

}
func (this *GoClass) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {

}

type GoFunc struct {
}

func (this *GoFunc) WriteHead(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)
	temp := `func ##class## ##head## `
	classfPrefix := ""
	if part.parent != nil && part.parent.partType == CodePart_class {
		last := strings.LastIndex(part.parent.data, "(")
		str := part.parent.data[0:last]
		str = strings.TrimSpace(str)
		classfPrefix = fmt.Sprintf("(self *%s)", str)
	}
	temp = strings.Replace(temp, "##class##", classfPrefix, 255)
	temp = strings.Replace(temp, "##head##", part.data, 255)
	temp = strings.Replace(temp, "(self,", "(", 255)
	temp = strings.Replace(temp, "(self)", "()", 255)
	buf.WriteString(temp)
}
func (this *GoFunc) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("{")
}
func (this *GoFunc) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)
	buf.WriteString("}")
}

type GoIf struct {
}

func (this *GoIf) WriteHead(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)

	data := part.data
	if strings.Contains(data, " in [") {
		fmt.Println("if in=>", data)
		//X in [-233, -232, -247]
		//a==b and value in [1, 2] and c==d
		s := strings.Split(data, " in [")
		if len(s) == 2 {
			fmt.Println("if in1=>", s[0], "=>", s[1])
			left, right := s[0], s[1] //a==b and value    =>     1, 2] and c==d
			llpos := strings.LastIndex(left, " ")
			var head string
			if llpos > 0 {
				head = left[0 : llpos+1]
				left = left[llpos+1:]
			}

			//fmt.Println("trans append2:", left, "=>", right)
			index := strings.LastIndex(right, "]")
			if index >= 0 {
				fmt.Println("if in2=>", right, index)

				var end = ""
				if index < len(right)-1 {
					end = right[index+1:] //and c==d
					fmt.Println("if end=>", end)
				}

				right = right[0:index] //1,2
				rights := strings.Split(right, ",")
				if len(rights) > 0 {
					fmt.Println("if in3=>", data)
					data = fmt.Sprintf("(%s==%s", left, rights[0])
					for i := 1; i < len(rights); i++ {
						data += fmt.Sprintf("||%s==%s", left, rights[i])
					}
					data += ")"
				}
				data = head + data + end
			}
		}
	}

	temp := `if ##head## `
	temp = strings.Replace(temp, "##head##", data, 255)
	buf.WriteString(temp)
}
func (this *GoIf) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("{")
}
func (this *GoIf) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)
	buf.WriteString("}")
}

type GoIfelse struct {
}

func (this *GoIfelse) WriteHead(part *CodePart, buf *bytes.Buffer) {
	//删掉一个\n
	//var left []byte
	//left = append(left, buf.Bytes()...)
	//buf.Reset()
	//buf.Write(left[0 : len(left)-1])

	temp := ` else if ##head## `
	temp = strings.Replace(temp, "##head##", part.data, 255)
	buf.WriteString(temp)
}
func (this *GoIfelse) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("{")
}
func (this *GoIfelse) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)
	buf.WriteString("}")
}

type GoElse struct {
}

func (this *GoElse) WriteHead(part *CodePart, buf *bytes.Buffer) {
	//删掉一个\n
	//var left []byte
	//left = append(left, buf.Bytes()...)
	//buf.Reset()
	//buf.Write(left[0 : len(left)-1])

	temp := ` else `
	buf.WriteString(temp)
}
func (this *GoElse) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("{")
}
func (this *GoElse) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)
	buf.WriteString("}")
}

type GoFor struct {
}

//for _guildId in list:								 slice
//for uid in guild.guildMember.iterkeys():           map.keys
//for guild in self.guildDict.itervalues():          map.values
//for uid, member in guild.guildMember.iteritems():  map k v
//for i in xrange(cardNum): 						 for i

func (this *GoFor) WriteHead(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)
	data := part.data

	if strings.HasSuffix(data, ".iteritems()") {
		//for uid, member in guild.guildMember.iteritems():  map k v
		data = strings.TrimSuffix(data, ".iteritems()") //remove tail
		data = strings.Replace(data, " in ", " :=range ", 1024)
	} else if strings.HasSuffix(data, ".itervalues()") {
		//for guild in self.guildDict.itervalues():          map.values
		s := strings.Split(data, " in ")
		head, tail := s[0], s[1]
		data = fmt.Sprintf("_,%s :=range %s", head, tail)
		data = strings.TrimSuffix(data, ".itervalues()") //remove tail
	} else if strings.HasSuffix(data, ".iterkeys()") {
		//for uid in guild.guildMember.iterkeys():           map.keys
		s := strings.Split(data, " in ")
		head, tail := s[0], s[1]
		data = fmt.Sprintf("%s,_ :=range %s", head, tail)
		data = strings.TrimSuffix(data, ".iterkeys()") //remove tail
	} else if strings.Contains(data, " xrange(") {
		//for i in xrange(cardNum): 						 for i
		s := strings.Split(data, " in xrange(")
		head, tail := s[0], s[1]
		tail = strings.TrimSuffix(tail, ")") //remove tail
		data = fmt.Sprintf("%s:=0;%s<%s;%s++", head, head, tail, head)
	} else {
		//for _guildId in list:								 slice
		data = "_," + data
		data = strings.Replace(data, " in ", " :=range ", 1024)
	}

	temp := `for ##head## `
	temp = strings.Replace(temp, "##head##", data, 1024)
	buf.WriteString(temp)

}
func (this *GoFor) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("{")
}
func (this *GoFor) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)
	buf.WriteString("}")
}

type GoRaw struct {
}

func (this *GoRaw) WriteHead(part *CodePart, buf *bytes.Buffer) {
	buf.WriteString("\n")
	part.writeScale(buf)

	data := part.data
	trim := strings.TrimSpace(data)
	if len(trim) > 0 && trim[0] != '#' { //不是注释

		if strings.Contains(data, ".append(") {
			//cardInfo.equipList.append(1)转换
			//fmt.Println("trans append:", data)
			s := strings.Split(data, ".append(")
			if len(s) >= 2 {
				left, right := s[0], s[1] //cardInfo.equipList,1)
				//fmt.Println("trans append2:", left, "=>", right)
				index := strings.LastIndex(right, ")")
				if index >= 0 {
					right = right[0:index] //1
					data = fmt.Sprintf("%s=append(%s,%s)", left, left, right)
				}
			}

		} else if strings.HasPrefix(trim, "del ") {
			//del self.buffList[i-1]
			trim = strings.TrimPrefix(trim, "del ")
			li := strings.Index(trim, "[")
			if li > 0 {
				left := trim[0:li]
				ri := strings.LastIndex(trim, "]")
				if ri > 0 {
					right := trim[li:ri]
					data = fmt.Sprintf("%s = append(%s[0:%s], %s[%s+1:]...)", left, left, right, left, right)
				}
			}
		} else if strings.HasPrefix(trim, "print ") {
			//print "layout"
			trim = strings.TrimPrefix(trim, "print ")
			data = fmt.Sprintf("fmt.Println(%s)", trim)

		}
	}

	buf.WriteString(data)
}
func (this *GoRaw) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {

}
func (this *GoRaw) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {
	//buf.WriteString("\n")
}
