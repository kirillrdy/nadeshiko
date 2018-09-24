package main

import (
	"net/http"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/jquery"
	"github.com/kirillrdy/web/html"
	"github.com/sparkymat/webdsl/css"
)

const websocketPath string = "/.websocket"

const click css.Class = "click"

func helloWorldHandler(document *nadeshiko.Document) {
	document.JQuery(css.Body).HTML(html.H1().Class("click").Text("Hello World !!!"))
	document.JQuery(click).Click(func() {
		document.JQuery(click).SetText("Oh now")
	})
}

func httpHandler(response http.ResponseWriter, request *http.Request) {
	nadeshiko.Page(websocketPath).WriteTo(response)
}

func main() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc(jquery.WebPath, jquery.FileHandler)
	http.Handle(websocketPath, nadeshiko.Handler(helloWorldHandler))
	http.ListenAndServe(":3000", nil)
}
