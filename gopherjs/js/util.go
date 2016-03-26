package js

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/kirillrdy/nadeshiko/html"
	"honnef.co/go/js/dom"
)

func SetTitle(title string) {
	js.Global.Get("document").Set("title", title)
}

func getBody() dom.Element {
	return dom.GetWindow().Document().QuerySelector("body")
}

func SetBody(content html.Node) {
	getBody().SetInnerHTML(content.String())
}

func NavigateTo(path string) {
	js.Global.Get("window").Get("history").Call("pushState", nil, nil, path)
	applyRoute()
}
