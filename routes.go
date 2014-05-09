package nadeshiko

import "net/http"

const GET = "GET"
const POST = "POST"

type Routes []Route

func (routes *Routes) Get(path string, handler HttpHander) {
	route := Route{path, GET, handler}
	*routes = append(*routes, route)
}

func (routes *Routes) Post(path string, handler HttpHander) {
	route := Route{path, POST, handler}
	*routes = append(*routes, route)
}

func (routes *Routes) Activity(path string, activity Activity) {
	// need to compose our own handler that servers index.html for nadeshiko
	handler := func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, nadeshikoPublicDir()+"/index.html")
	}

	//TODO hacky, we should get rid of concept of default activity
	DefaultActivity = activity

	route := Route{path, GET, handler}
	*routes = append(*routes, route)
}
