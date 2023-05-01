package genworld

//go:generate ./gotext -srclang=en update -out=catalog/catalog.go -lang=en,zh

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Proc(lang language.Tag) {
	p := message.NewPrinter(lang)

	p.Printf("Hello!")
}
