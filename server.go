package nadeshiko

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
)

const NADESHIKO_VERSION = "0.1.0"

type OverSocketCallback struct {
	connection  *Connection
	OneTimeOnly bool
	Callback    func(...string)
}

var Callbacks map[string]OverSocketCallback
var Verbose bool

// Serves base.html, thin html client
func fileServer(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("\n\n============================================\n")
	requested_path := req.RequestURI

	if requested_path == "/" {
		requested_path = "/base.html"
	}

	fmt.Printf("GET: %q \n", requested_path)
	fmt.Printf("User-Agent: %s \n", req.Header["User-Agent"])

	w.Header().Set("Server", "Nadeshiko "+NADESHIKO_VERSION)

	http.ServeFile(w, req, "public"+requested_path)

}

func Start(activity Activity, port int, verbose bool) {

	DefaultActivity = activity

	//TODO Perhaps we dont need to export these
	Callbacks = make(map[string]OverSocketCallback)
	Notifications = make(map[string][]*Connection)
	Verbose = verbose

	http.Handle("/websocket_client", websocket.Handler(websocketServer))
	http.HandleFunc("/", fileServer)

	fmt.Println("Started Nadeshiko Server " + NADESHIKO_VERSION)
	fmt.Printf("Listening http://localhost:%d/\n", port)

	listenOn := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listenOn, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
