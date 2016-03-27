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
	GetDocument().SetTitle("Page Not Found")
	GetDocument().GetBody().SetInnerHTML(
		html.H1().Text("ERROR Page not found"),
	)
}

func addEventListener(object *js.Object, eventName string, callback func(*js.Object)) {
	object.Call("addEventListener", eventName, callback)
}

func handlePopState() {
	popStateCallBack := func(event *js.Object) {
		applyRoute()
	}
	addEventListener(js.Global.Get("window"), "popstate", popStateCallBack)
}

func applyRoute() {
	handler := findRouteHandler(currentPath())
	handler()
}

//TODO rename to something better
func RouterRun() {
	handlePopState()
	addEventListener(js.Global.Get("document"), "DOMContentLoaded", func(event *js.Object) {
		applyRoute()
	})

}

//NavigateTo changes current path and calls appropriate route handler
//TODO move this to router
func NavigateTo(path string) {
	js.Global.Get("window").Get("history").Call("pushState", nil, nil, path)
	applyRoute()
}
