// +build js

package main

import (
	"github.com/kirillrdy/nadeshiko/gopherjs/js"
	"github.com/kirillrdy/nadeshiko/html"
	"honnef.co/go/js/dom"
)

func mainPage(document dom.Document) {
	js.SetTitle("Main Page")
	js.SetBody(html.Div().Children(
		html.H1().Text("Hello world"),
	))
}

func main() {
	js.AddRoute("/", mainPage)
	js.RouterRun()
}
