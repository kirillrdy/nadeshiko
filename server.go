package nadeshiko

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	//"io"
	//"io/ioutil"
	"net/http"



)

const NADESHIKO_VERSION = "0.1.0"

type OverSocketCallback struct {
	ws	*WebsocketConnection
	OneTimeOnly bool
	Callback func(...string)

}

var	Callbacks map[string] OverSocketCallback
var verbose bool


type WebsocketConnection websocket.Conn

func (ws *WebsocketConnection) SendMessage(message string) {
	real_websocket := websocket.Conn(*ws)
	websocket.Message.Send(&real_websocket, message)


	// TODO perhaps good to keep, on Send cleanup sockets
	// however most of those should be caught in the main loop
	//err := websocket.Message.Send(&real_websocket, message)

	//if err != nil {
	//	for k, v := range Notifications {
	//		var new_list []JFunc
	//		for _, jFunction := range v {
	//			if jFunction("").ws != ws {
	//				new_list = append(new_list,jFunction)
	//			}
	//		}
	//		Notifications[k] = new_list
	//	}

	//	for callback_id, callbackStruct := range Callbacks {
	//		if callbackStruct.ws == ws {
	//			delete(Callbacks, callback_id)
	//			fmt.Printf("Removing callback %s for disconnected client\n",callback_id)
	//		}
	//	}
	//}



	//TODO if we fail send we should remove Callbacks and Notifications
	if verbose {
		fmt.Printf("send: %s\n", message)
	}
}



// Serves base.html, thin html client
func fileServer(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("\n\n============================================\n")
	requested_path := req.RequestURI

	if requested_path == "/" {
		requested_path = "/base.html"
	}

	fmt.Printf("GET: %q \n", requested_path)
	fmt.Printf("User-Agent: %s \n",req.Header["User-Agent"])

	w.Header().Set("Server","Nadeshiko "+ NADESHIKO_VERSION)

	http.ServeFile(w, req, "public" + requested_path)

}


func Start(port int) {

	//TODO Perhaps we dont need to export these
	Callbacks = make(map[string] OverSocketCallback)
	Notifications = make(map[string] []WebsocketConnection)
	verbose = false

	http.Handle("/websocket_client", websocket.Handler(websocketServer))
	http.HandleFunc("/", fileServer)

	fmt.Println("Started Nadeshiko Server " + NADESHIKO_VERSION)
	fmt.Printf("Listening http://localhost:%d/\n",port)

	listenOn := fmt.Sprintf(":%d",port)
	err := http.ListenAndServe(listenOn, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
