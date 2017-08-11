package main

import "bytes"

type ITranslate interface {
	WriteHead(*CodePart, *bytes.Buffer)
	WriteBodyPre(*CodePart, *bytes.Buffer)
	WriteBodyEnd(*CodePart, *bytes.Buffer)
}

func GetGolangTrans(ctype CodePartType) ITranslate {
	switch ctype {
	case CodePart_class:
		return new(GoClass)
	case CodePart_func:
		return new(GoFunc)
	case CodePart_for:
		return new(GoFor)
	case CodePart_if:
		return new(GoIf)
	case CodePart_elseif:
		return new(GoIfelse)
	case CodePart_else:
		return new(GoElse)
	case CodePart_raw:
		return new(GoRaw)
	default:
		return new(GoRaw)
	}
}
