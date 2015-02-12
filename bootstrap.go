package nadeshiko

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

//TODO get rid of all of this
func (document *Document) IncludeBootstrapCdn() {
	document.JQuery(css.Body).Append(html.Link().Rel("stylesheet").Href("//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css"))
}

//TODO get rid of all of this
func (document *Document) IncludeBootstrapJsCdn() {
	document.JQuery(css.Body).Append(html.Script().Src("//netdna.bootstrapcdn.com/bootstrap/3.0.0/js/bootstrap.min.js"))
}

//TODO get rid of all of this
func (element jQuerySelector) Tooltip() {
	element.zeroArgumentMethod("tooltip")
}
