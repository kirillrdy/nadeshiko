package main

import (
	"io"
	"net/http"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
)

func handler(response http.ResponseWriter, request *http.Request) {
	page := html.Html().Children(
		html.Head().Children(
			html.Title().Text("Todo App"),
		),
		html.Body().Children(
			html.H1().Text("Hello World"),
		),
	)
	io.WriteString(response, page.String())
}

func main() {
	nadeshiko.Get("/", handler)
	nadeshiko.Start()
}
