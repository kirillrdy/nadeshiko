package main

import (
	"io"
	"net/http"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

func handler(document *nadeshiko.Document) {
	document.JQuery(css.Body).Append(html.H1().Text("Hello World !!!"))
}

func httpHandler(response http.ResponseWriter, request *http.Request) {

	page := html.Html().Children(
		html.Head().Children(
			nadeshiko.NadeshikoScripts()...,
		),
		html.Body(),
	)
	io.WriteString(response, page.String())

}

func main() {
	http.HandleFunc("/", httpHandler)

	//XXX for now just have path + .websocket pattern
	http.Handle("/.websocket", nadeshiko.Handler(handler))
	http.ListenAndServe(":3000", nil)
}
