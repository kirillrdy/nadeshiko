//Nadeshiko is my toy library for building interesting web apps
package nadeshiko

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

const NADESHIKO_VERSION = "0.1.0"

var defaultRoutes routes

func Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	defaultRoutes.Get(path, handler)
}

func Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	defaultRoutes.Post(path, handler)
}

func Nadeshiko(path string, handler func(*Document)) {
	defaultRoutes.Nadeshiko(path, handler)
}

func WebSocket(path string, handler func(*Document)) {
	defaultRoutes.WebSocket(path, handler)
}

func startWithPortVerbose(port int, verbose bool) {

	log.Println("Starting Nadeshiko Server " + NADESHIKO_VERSION)

	listenOn := fmt.Sprintf(":%d", port)

	log.Printf("Listening http://0.0.0.0:%d/\n", port)
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

func Start() {
	startWithPortVerbose(3000, false)
}
