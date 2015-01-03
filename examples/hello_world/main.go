package main

import (
	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

func handler(document *nadeshiko.Document) {
	document.JQuery(css.Body).Html(html.H1().Text("Hello World !!!"))
}

func main() {
	nadeshiko.Nadeshiko("/", handler)
	nadeshiko.Start()
}
