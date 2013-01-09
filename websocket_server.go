package nadeshiko

import (
	"code.google.com/p/go.net/websocket"
	"strconv"
	"encoding/json"
	"log"
	"io"
)



func websocketServer(ws *websocket.Conn) {
	log.Printf("New client connection on %#v\n", &ws)

	socket := WebsocketConnection(*ws)
	connection := Connection{websocket: &socket}

	DefaultActivity.Start(&connection)


	for {
		var buf string

		//Consider using JSON codec for websocket
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

		if callbackStruct, ok := callbacks[json_array[0]]; ok {
			callbackStruct.callback(json_array...)
			if callbackStruct.oneTimeOnly {
				if Verbose {
					log.Printf("Removing one-time callback \n")
				}
				deleteCallback <- json_array[0]
			}
			if Verbose {
				log.Printf("Current callbacks count %d \n", len(callbacks))
			}

		} else {
			log.Printf("Cant find callback for %s \n", json_array[0])
		}
	}

	CleanupNotification <- &connection

	for callback_id, callbackStruct := range callbacks {
		if callbackStruct.connection == &connection {
			if Verbose {
				log.Printf("Removing callback %s for disconnected client\n", callback_id)
			}
			deleteCallback <- callback_id
		}
	}
	log.Printf("Current callbacks count %d \n", len(callbacks))

}
