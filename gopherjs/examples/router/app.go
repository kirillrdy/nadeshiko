// +build js

package main

import (
	"github.com/kirillrdy/nadeshiko/gopherjs/js"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
	"honnef.co/go/js/dom"
)

var paths = struct {
	main   string
	second string
}{
	"/",
	"/second",
}

var h1Id css.Id = "my-h1"

func mainPage(document dom.Document) {

	js.SetTitle("Main Page")
	js.SetBody(html.Div().Children(
		html.H1().Id(h1Id).Text("Click Me"),
		html.A().Href(paths.second).Text("Second page"),
	))

	h1 := document.QuerySelector(h1Id.Selector())

	h1.AddEventListener("click", true, func(event dom.Event) {
		js.NavigateTo(paths.second)
	})
}

func secondPage(document dom.Document) {
	js.SetTitle("Second Page")
	js.SetBody(
		html.Div().Children(
			html.H1().Id(h1Id).Text("I am second page"),
			html.A().Href(paths.main).Text("First page"),
		),
	)

	h1 := document.QuerySelector(h1Id.Selector())

	h1.AddEventListener("click", true, func(event dom.Event) {
		js.NavigateTo(paths.main)
	})
}

func main() {
	js.AddRoute(paths.main, mainPage)
	js.AddRoute(paths.second, secondPage)
	js.RouterRun()
}
