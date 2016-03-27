package js

import (
	gopherjs "github.com/gopherjs/gopherjs/js"
	"github.com/sparkymat/webdsl/css"
)

//Document represents browsers document
type Document struct {
	object *gopherjs.Object
}

//QuerySelector equivaltent to javascripts querySelector
func (document Document) QuerySelector(selector css.Selector) Element {
	object := document.object.Call("querySelector", selector.Selector())
	return Element{object: object}
}

//SetTitle sets the title of the document
func (document Document) SetTitle(title string) {
	document.object.Set("title", title)
}

func (document Document) GetBody() Element {
	return document.QuerySelector(css.Body)
}
