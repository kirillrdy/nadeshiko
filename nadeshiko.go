package nadeshiko

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
	"log"
	"fmt"
)

const NADESHIKO_VERSION = "0.1.0"

var DefaultActivity Activity

func StartWithPortVerbose(activity Activity, port int, verbose bool) {
	DefaultActivity = activity
	Verbose = verbose

	log.Println("Starting Nadeshiko Server " + NADESHIKO_VERSION)
	http.Handle("/nadeshiko_socket", websocket.Handler(websocketServer))
	http.HandleFunc("/", fileServer)

	listenOn := fmt.Sprintf(":%d", port)

	log.Printf("Listening http://localhost:%d/\n", port)
	err := http.ListenAndServe(listenOn, nil)
	if err != nil {
		log.Fatalln("ListenAndServe: " + err.Error())
	}

}

func StatWithPort(activity Activity, port int) {
	StartWithPortVerbose(activity, port, false)
}

func Start(activity Activity) {
	StartWithPortVerbose(activity, 3000, false)
}
