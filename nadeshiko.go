package nadeshiko

import (
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

const NADESHIKO_VERSION = "0.1.0"

var DefaultActivity Activity

func startWithPortVerbose(routes Routes, port int, verbose bool) {
	//DefaultActivity = activity
	Verbose = verbose

	log.Println("Starting Nadeshiko Server " + NADESHIKO_VERSION)

	//TODO hack, to restore our websocket handler
	routes.Get("/nadeshiko_socket", websocket.Handler(websocketServer).ServeHTTP)

	listenOn := fmt.Sprintf(":%d", port)

	log.Printf("Listening http://0.0.0.0:%d/\n", port)
	err := http.ListenAndServe(listenOn, httpHandler{routes: routes})
	if err != nil {
		log.Fatalln("ListenAndServe: " + err.Error())
	}
}

func StartWithPortVerbose(routes Routes, port int) {
	startWithPortVerbose(routes, port, true)
}

func StartWithPort(routes Routes, port int) {
	startWithPortVerbose(routes, port, false)
}

func Start(routes Routes) {
	startWithPortVerbose(routes, 3000, false)
}

func StartActivity(activity Activity) {
	routes := Routes{}
	routes.Activity("/", activity)
	startWithPortVerbose(routes, 3000, false)
}
