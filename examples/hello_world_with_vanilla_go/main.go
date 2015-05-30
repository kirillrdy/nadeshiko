package main

import (
	"net/http"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/nadeshiko/jquery"
	"github.com/sparkymat/webdsl/css"
)

const rootPath string = "/"
const websocketPath string = "/.websocket"

func helloWorldHandler(document *nadeshiko.Document) {
	document.JQuery(css.Body).HTML(html.H1().Text("Hello World !!!"))
}

func httpHandler(response http.ResponseWriter, request *http.Request) {
	nadeshiko.Page(websocketPath).WriteTo(response)
}

func main() {
	http.DefaultServeMux.HandleFunc(rootPath, httpHandler)
	http.DefaultServeMux.HandleFunc(jquery.WebPath, jquery.FileHandler)
	http.DefaultServeMux.Handle(websocketPath, nadeshiko.Handler(helloWorldHandler))
	http.ListenAndServe(":3000", nil)
}
