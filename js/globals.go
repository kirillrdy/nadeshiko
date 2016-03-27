package js

import (
	gopherjs "github.com/gopherjs/gopherjs/js"
)

func GetDocument() Document {
	object := gopherjs.Global.Get("document")
	return Document{object: object}
}
