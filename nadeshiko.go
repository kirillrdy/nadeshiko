package nadeshiko

import (
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

const NADESHIKO_VERSION = "0.1.0"

var DefaultActivity Activity

func startWithPortVerbose(port int, verbose bool) {
	//DefaultActivity = activity
	Verbose = verbose

	log.Println("Starting Nadeshiko Server " + NADESHIKO_VERSION)

	//TODO hack, to restore our websocket handler
	defaultRoutes.Get("/nadeshiko_socket", websocket.Handler(websocketServer).ServeHTTP)

	listenOn := fmt.Sprintf(":%d", port)

	log.Printf("Listening http://0.0.0.0:%d/\n", port)
	err := http.ListenAndServe(listenOn, httpHandler{routes: defaultRoutes})
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
