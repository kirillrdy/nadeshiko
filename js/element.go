package js

import (
	gopherjs "github.com/gopherjs/gopherjs/js"
	"github.com/kirillrdy/web/html"
)

type Element struct {
	object *gopherjs.Object
}

func (element Element) AddEventListener(eventName string, callback func(*gopherjs.Object)) {
	element.object.Call("addEventListener", eventName, callback)
}

//SetBody sets the content of the body, in the future this should be replaced by own dom wrapping
func (element Element) SetInnerHTML(content html.Node) {
	element.object.Set("innerHTML", content.String())
}
