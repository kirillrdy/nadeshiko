//Nadeshiko is my toy library for building interesting web apps
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

func Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	defaultRoutes.get(path, handler)
}

func Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	defaultRoutes.post(path, handler)
}

func Nadeshiko(path string, handler func(*Document)) {
	defaultRoutes.nadeshiko(path, handler)
}

func WebSocket(path string, handler func(*Document)) {
	defaultRoutes.webSocket(path, handler)
}

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

func StartWithPortVerbose(port int) {
	startWithPortVerbose(port, true)
}

func StartWithPort(port int) {
	startWithPortVerbose(port, false)
}

//Starts server in non verbose mode and on port 3000
func Start() {
	startWithPortVerbose(3000, false)
}

//TODO move somewhere else
func NadeshikoScripts() []html.Node {
	return []html.Node{
		html.Script().Attribute("src", jquery.WebPath),
		html.Script().Attribute("src", NadeshikoJsWebPath),
	}
}

const NadeshikoJsWebPath = "/socket_init.js"

func NadeshikoJsHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, nadeshikoPublicDir()+NadeshikoJsWebPath)
}
