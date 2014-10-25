package main

import (
	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
)

func handler(document *nadeshiko.Document) {
	document.JQuery("body").Append(html.H1().Text("Hello World !!!").String())
}

func main() {
	nadeshiko.Nadeshiko("/", handler)
	nadeshiko.Start()
}
