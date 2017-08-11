package main

import (
	"bytes"
	"fmt"
	"strings"
)

type GoClass struct {
}

func (this *GoClass) WriteHead(part *CodePart, buf *bytes.Buffer) {
	last := strings.LastIndex(part.data, "(")
	str := part.data[0:last]
	//fmt.Println("str:%d %d %s", last, len(part.data), str)
	temp := "type ##head## struct {}\n"
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
	if part.parent != nil {
		last := strings.LastIndex(part.parent.data, "(")
		str := part.parent.data[0:last]
		str = strings.TrimSpace(str)
		classfPrefix = fmt.Sprintf("(self *%s)", str)
	}
	temp = strings.Replace(temp, "##class##", classfPrefix, 255)
	temp = strings.Replace(temp, "##head##", part.data, 255)
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
	temp := `if ##head## `
	temp = strings.Replace(temp, "##head##", part.data, 255)
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
	buf.WriteString(part.data)
}
func (this *GoRaw) WriteBodyPre(part *CodePart, buf *bytes.Buffer) {

}
func (this *GoRaw) WriteBodyEnd(part *CodePart, buf *bytes.Buffer) {
	//buf.WriteString("\n")
}
