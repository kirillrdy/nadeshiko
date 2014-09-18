package nadeshiko

import (
	"encoding/json"
	"io"
	"log"
	"strconv"

	"code.google.com/p/go.net/websocket"
)

func websocketServer(handler func(*Document)) func(*websocket.Conn) {
	return func(connection *websocket.Conn) {
		log.Printf("New client connection\n")

		document := Document{websocket: connection}
		handler(&document)

		for {
			var buf string

			//TODO Consider using JSON codec for websocket
			err := websocket.Message.Receive(connection, &buf)
			if err != nil {
				if err == io.EOF {
					log.Println("Websocket client disconnected")
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
					deleteCallback(json_array[0])
				}
				if Verbose {
					log.Printf("Current callbacks count %d \n", len(callbacks))
				}

			} else {
				log.Printf("Can't find callback for %s \n", json_array[0])
			}
		}

		for callback_id, callbackStruct := range callbacks {
			if callbackStruct.connection == nil {
				log.Fatalln("This should not happen")
			}
			if callbackStruct.connection == connection {
				if Verbose {
					log.Printf("Removing callback %s for disconnected client\n", callback_id)
				}
				deleteCallback(callback_id)
			}
		}

		if Verbose {
			log.Printf("Current callbacks count %d \n", len(callbacks))
		}
	}
}
