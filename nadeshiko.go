package nadeshiko

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
	"log"
	"fmt"
)

//TODO consider getting rid of this
var DefaultActivity Activity

func Start(activity Activity, port int, verbose bool) {

	DefaultActivity = activity


	Verbose = verbose

	http.Handle("/websocket_client", websocket.Handler(websocketServer))
	http.HandleFunc("/", fileServer)

	log.Println("Started Nadeshiko Server " + NADESHIKO_VERSION)
	log.Printf("Listening http://localhost:%d/\n", port)

	listenOn := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listenOn, nil)
	if err != nil {
		log.Fatalln("ListenAndServe: " + err.Error())
	}
}
