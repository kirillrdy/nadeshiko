// +build js

package main

import (
	gopherjs "github.com/gopherjs/gopherjs/js"
	"github.com/kirillrdy/nadeshiko/js"
	"github.com/kirillrdy/web/html"
	"github.com/sparkymat/webdsl/css"
)

var paths = struct {
	main   string
	second string
}{
	"/",
	"/second",
}

var h1Id css.Id = "my-h1"

func mainPage() {

	js.GetDocument().SetTitle("Main Page")
	js.GetDocument().GetBody().SetInnerHTML(html.Div().Children(
		html.H1().Id(h1Id).Text("Click Me"),
		html.A().Href(paths.second).Text("Second page"),
	))

	navUsingH1(paths.second)

}

func secondPage() {
	js.GetDocument().SetTitle("Second Page")
	js.GetDocument().GetBody().SetInnerHTML(
		html.Div().Children(
			html.H1().Id(h1Id).Text("I am second page"),
			html.A().Href(paths.main).Text("First page"),
		),
	)

	navUsingH1(paths.main)

}

func navUsingH1(path string) {

	//TODO make this better
	h1 := js.GetDocument().QuerySelector(h1Id)
	h1.AddEventListener("click", func(event *gopherjs.Object) {
		js.NavigateTo(path)
	})
}

func main() {
	js.AddRoute(paths.main, mainPage)
	js.AddRoute(paths.second, secondPage)
	js.RouterRun()
}
