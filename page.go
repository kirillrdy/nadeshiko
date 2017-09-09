package nadeshiko

import (
	"io"

	"github.com/kirillrdy/web/html"
)

//This is the default blank html page. that contains required clientside js code to have nadeshiko working
func Page(path string) io.WriterTo {
	return html.Html().Children(
		html.Head().Children(
			Scripts(path)...,
		),
		html.Body(),
	)
}
