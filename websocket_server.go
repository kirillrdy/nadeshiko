package nadeshiko

import "code.google.com/p/go.net/websocket"
import "strconv"
import "encoding/json"
import "log"
import "io"

func websocketServer(ws *websocket.Conn) {
	log.Printf("New client connection on %#v\n", &ws)

	socket := WebsocketConnection(*ws)
	connection := Connection{websocket: &socket}

	DefaultActivity.Start(&connection)

	//TODO fix this testing thing
	//if *test_env {
	//	go func(){
	//		ApplicationTest(j)
	//	}()
	//}

	for {
		var buf string
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Client Disconnected")
			} else {
				log.Printf("ERROR reading from socket: %v \n", err)
			}
			break
		}

		if Verbose {
			quoted_contet := strconv.Quote(buf)
			log.Printf("received: %s\n", quoted_contet)
		}

		var json_array []string
		json.Unmarshal([]byte(buf), &json_array)

		//This if statment not really needed,
		// since websocket.Message.Receive should catch most of errors
		// but good to keep for debugging
		if callbackStruct, ok := Callbacks[json_array[0]]; ok {
			callbackStruct.Callback(json_array...)
			if callbackStruct.OneTimeOnly {
				if Verbose {
					log.Printf("Removing one-time callback, curent count %d \n", len(Callbacks))
				}
				delete(Callbacks, json_array[0])
			}
			if Verbose {
				log.Printf("Current callbacks count %d \n", len(Callbacks))
			}

		} else {
			log.Printf("Cant find callback for %s \n", json_array[0])
		}
	}

	CleanupNotification(&connection)

	for callback_id, callbackStruct := range Callbacks {
		if callbackStruct.connection == &connection {
			delete(Callbacks, callback_id)
			if Verbose {
				log.Printf("Removing callback %s for disconnected client\n", callback_id)
			}
		}
	}
	log.Printf("Current callbacks count %d \n", len(Callbacks))

}
