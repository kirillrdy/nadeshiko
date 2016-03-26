package js

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/kirillrdy/nadeshiko/html"
)

func currentPath() string {
	return js.Global.Get("document").Get("location").Get("pathname").String()
}

type route struct {
	path    string
	handler func()
}

var routes []route

func AddRoute(path string, handler func()) {
	routes = append(routes, route{path: path, handler: handler})
}

func findRouteHandler(path string) func() {
	for _, route := range routes {
		//TODO this needs to be much better
		if route.path == path {
			return route.handler
		}
	}
	return pageNotFound
}

func pageNotFound() {
	SetTitle("Page Not Found")
	SetBody(
		html.H1().Text("ERROR Page not found"),
	)
}

func handlePopState() {
	js.Global.Get("window").Call("addEventListener", "popstate", func(event js.Object) {
		applyRoute()
	})
}

func applyRoute() {
	handler := findRouteHandler(currentPath())
	handler()
}

//TODO rename to something better
func RouterRun() {
	//TODO for some reasong this needs to be done before any call to dom
	js.Global.Get("document").Call("write", "Init Router")

	handlePopState()
	applyRoute()
}
