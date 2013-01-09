package nadeshiko

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
	"log"
	"fmt"
)

var DefaultActivity Activity

func Start(activity Activity, port int, verbose bool) {

	DefaultActivity = activity
	Verbose = verbose

	log.Println("Starting Nadeshiko Server " + NADESHIKO_VERSION)
	http.Handle("/websocket_client", websocket.Handler(websocketServer))
	http.HandleFunc("/", fileServer)

	listenOn := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listenOn, nil)
	if err != nil {
		log.Fatalln("ListenAndServe: " + err.Error())
	}

	log.Printf("Listening http://localhost:%d/\n", port)
}
