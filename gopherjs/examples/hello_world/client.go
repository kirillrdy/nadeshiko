// +build js

package main

import (
	"github.com/kirillrdy/nadeshiko/js"
	"github.com/kirillrdy/web/html"
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
