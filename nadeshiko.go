//Package nadeshiko is my toy library for building interesting web apps
/*
Examples

please see github.com/kirillrdy/nadeshiko/examples package

*/
package nadeshiko

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/nadeshiko/jquery"
)

//TODO need better versioning method
const Version = "0.1.0"

var defaultRoutes routes

//Get adds handler to a GET request
func Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	defaultRoutes.get(path, handler)
}

//Post adds handler to a POST request
func Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	defaultRoutes.post(path, handler)
}

//Nadeshiko adds nadeshiko handler to a given path
func Nadeshiko(path string, handler func(*Document)) {
	defaultRoutes.nadeshiko(path, handler)
}

//TODO Do we actually need this ? should this be internal
//func WebSocket(path string, handler func(*Document)) {
//	defaultRoutes.webSocket(path, handler)
//}

func startWithPortVerbose(port int, verbose bool) {

	log.Println("Starting Nadeshiko Server " + Version)

	listenOn := fmt.Sprintf(":%d", port)

	log.Printf("Listening http://0.0.0.0:%d/\n", port)

	//TODO find better place to initialise defaultRoutes or get rid of defaultRoutes
	defaultRoutes.get(jquery.WebPath, jquery.FileHandler)

	err := http.ListenAndServe(listenOn, context.ClearHandler(internalHttpHandler{routes: defaultRoutes}))
	if err != nil {
		log.Fatalln("ListenAndServe: " + err.Error())
	}
}

//StartWithPortVerbose starts server on a given port in verbose mode
func StartWithPortVerbose(port int) {
	startWithPortVerbose(port, true)
}

//StartWithPort starts server on a given port in non verbose mode
func StartWithPort(port int) {
	startWithPortVerbose(port, false)
}

//Start starts server in non verbose mode and on port 3000
func Start() {
	startWithPortVerbose(3000, false)
}

//Scripts returns html.Nodes that are need to be placed on the page for nadeshiko to work
//TODO move somewhere else
func Scripts() []html.Node {
	return []html.Node{
		html.Script().Attribute("src", jquery.WebPath),
		html.Script().Attribute("src", NadeshikoJsWebPath),
	}
}

// Where to serve nadeshiko javascript file from
const NadeshikoJsWebPath = "/socket_init.js"

//JsHandler serves javascript file requires for having nadeshiko running
func JsHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, nadeshikoPublicDir()+NadeshikoJsWebPath)
}
