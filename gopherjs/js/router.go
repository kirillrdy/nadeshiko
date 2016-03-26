package js

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/kirillrdy/nadeshiko/html"
	"honnef.co/go/js/dom"
)

func currentPath() string {
	return js.Global.Get("document").Get("location").Get("pathname").String()
}

type route struct {
	path    string
	handler func(dom.Document)
}

var routes []route

func AddRoute(path string, handler func(dom.Document)) {
	routes = append(routes, route{path: path, handler: handler})
}

func findRoute(path string) func(dom.Document) {
	for _, route := range routes {
		//TODO this needs to be much better
		if route.path == path {
			return route.handler
		}
	}
	return pageNotFound
}

func pageNotFound(document dom.Document) {
	SetTitle("Page Not Found")
	SetBody(
		html.H1().Text("ERROR Page not found"),
	)
}

func handlePopState() {
	dom.GetWindow().AddEventListener("popstate", true, func(event dom.Event) {
		applyRoute()

	})
}

func applyRoute() {
	document := dom.GetWindow().Document()
	handler := findRoute(currentPath())
	handler(document)
}

//TODO rename to something better
func RouterRun() {
	//TODO for some reasong this needs to be done before any call to dom
	js.Global.Get("document").Call("write", "a")

	handlePopState()
	applyRoute()
}
