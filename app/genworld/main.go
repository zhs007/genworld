package main

import (
	"github.com/zhs007/genworld"
	"github.com/zhs007/goutils"
	"golang.org/x/text/language"
)

func main() {
	goutils.InitLogger("genworld", genworld.Version, "info", true, "./logs")

	genworld.Proc(language.English)
}
