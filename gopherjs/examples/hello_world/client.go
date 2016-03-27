// +build js

package main

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/nadeshiko/js"
)

func mainPage() {
	js.GetDocument().SetTitle("Main Page")
	js.GetDocument().GetBody().SetInnerHTML(html.Div().Children(
		html.H1().Text("Hello world"),
	))
}

func main() {
	js.AddRoute("/", mainPage)
	js.RouterRun()
}
