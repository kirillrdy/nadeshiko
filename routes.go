package nadeshiko

import (
	"net/http"

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

func (routes *routes) Nadeshiko(path string, handler func(Document)) {
	// need to compose our own handler that servers index.html for nadeshiko
	httpHandler := func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, nadeshikoPublicDir()+"/index.html")
	}

	customeWebsocketServer := websocketServer(handler)

	routes.Get(path+".websocket", websocket.Handler(customeWebsocketServer).ServeHTTP)

	routes.Get(path, httpHandler)
}
