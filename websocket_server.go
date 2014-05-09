package nadeshiko

import (
	"encoding/json"
	"io"
	"log"
	"strconv"

	"code.google.com/p/go.net/websocket"
)

func websocketServer(ws *websocket.Conn) {
	log.Printf("New client connection\n")

	socket := WebsocketConnection(*ws)
	connection := Connection{websocket: &socket}

	DefaultActivity.Start(&connection)

	for {
		var buf string

		//TODO Consider using JSON codec for websocket
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Webscoket client disconnected")
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
			if callbackStruct.oneTime {
				if Verbose {
					log.Printf("Removing one-time callback \n")
				}
				deleteCallback <- json_array[0]
			}
			if Verbose {
				log.Printf("Current callbacks count %d \n", len(callbacks))
			}

		} else {
			log.Printf("Can't find callback for %s \n", json_array[0])
		}
	}

	CleanupEventHandlers <- &connection

	for callback_id, callbackStruct := range callbacks {
		if callbackStruct.connection == &connection {
			if Verbose {
				log.Printf("Removing callback %s for disconnected client\n", callback_id)
			}
			deleteCallback <- callback_id
		}
	}

	if Verbose {
		log.Printf("Current callbacks count %d \n", len(callbacks))
	}
}
