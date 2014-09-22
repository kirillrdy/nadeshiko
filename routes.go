package nadeshiko

import (
	"io"
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"

	"code.google.com/p/go.net/websocket"
)

const GET = "GET"
const POST = "POST"

type routes []Route

func (routes *routes) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := Route{path, GET, handler}
	*routes = append(*routes, route)
}

func (routes *routes) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := Route{path, POST, handler}
	*routes = append(*routes, route)
}

func (routes *routes) Nadeshiko(path string, handler func(*Document)) {
	httpHandler := func(response http.ResponseWriter, request *http.Request) {
		page := html.Html().Children(
			html.Head().Children(
				html.Script().Attribute("src", ""),
				html.Script().Attribute("src", "/jquery-2.0.3.min.js"),
				html.Script().Attribute("src", "/jquery-ui-1.8.21.custom.min.js"),
				html.Script().Attribute("src", "/socket_init.js"),
			),
			html.Body(),
		)
		io.WriteString(response, page.String())
	}

	customeWebsocketServer := websocketServer(handler)

	routes.Get(path+".websocket", websocket.Handler(customeWebsocketServer).ServeHTTP)

	routes.Get(path, httpHandler)
}
