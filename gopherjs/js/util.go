package js

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/kirillrdy/nadeshiko/html"
)

//SetTitle sets the title of the document
func SetTitle(title string) {
	js.Global.Get("document").Set("title", title)
}

func getDocument() *js.Object {
	return js.Global.Get("document")
}

func getBody() *js.Object {
	return getDocument().Call("querySelector", "body")
}

//SetBody sets the content of the body, in the future this should be replaced by own dom wrapping
func SetBody(content html.Node) {
	//TODO replace with own dom thingy
	body := getBody()
	body.Set("innerHTML", content.String())
}

//NavigateTo changes current path and calls appropriate route handler
//TODO move this to router
func NavigateTo(path string) {
	js.Global.Get("window").Get("history").Call("pushState", nil, nil, path)
	applyRoute()
}
