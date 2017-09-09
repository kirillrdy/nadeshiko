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

	"github.com/kirillrdy/nadeshiko/jquery"
	"github.com/kirillrdy/web/html"
	"golang.org/x/net/websocket"
)

//TODO need better versioning method
const Version = "0.1.0"

func startWithPortVerbose(port int, verbose bool) {

	log.Println("Starting Nadeshiko Server " + Version)

	listenOn := fmt.Sprintf(":%d", port)

	log.Printf("Listening http://0.0.0.0:%d/\n", port)

	//TODO find better place to initialise defaultRoutes or get rid of defaultRoutes
	http.HandleFunc(jquery.WebPath, jquery.FileHandler)

	//TODO get rid of gorilla, rid of own router
	err := http.ListenAndServe(listenOn, nil)
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
// path argument is for where websocket handler will try to connect
func Scripts(path string) []html.Node {
	return []html.Node{
		html.Script().Attribute("src", jquery.WebPath),
		html.Script().TextUnsafe(SocketJSCode(path)),
	}
}

// Handler returns returns http.Handler for given nadeshiko handler so that it can be used with router of your choice
func Handler(nadeshikoHandler func(*Document)) http.Handler {
	return websocket.Handler(websocketServer(nadeshikoHandler))
}
