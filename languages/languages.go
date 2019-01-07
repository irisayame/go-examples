package main

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,ja-JP

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "github.com/xinnige/go-examples/languages/catalog"
)

func main() {
	p := message.NewPrinter(language.MustParse("ja-JP"))

	p.Printf("Hello world!")

	p.Println()

	booknum := 1500
	p.Printf("There are %v books in the library.", booknum)
	p.Println()
}
