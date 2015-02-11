package nadeshiko

import (
	"encoding/json"
	"io"
	"log"
	"strconv"

	"code.google.com/p/go.net/websocket"
)

//TODO this function is too long needs clean up
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

			if verbose {
				quotedContet := strconv.Quote(buf)
				log.Printf("received: %s\n", quotedContet)
			}

			var jsonArray []string
			json.Unmarshal([]byte(buf), &jsonArray)

			//TODO check length of jsonArray
			if callbackStruct, ok := callbacks[jsonArray[0]]; ok {
				//TODO danger, huge possibility for data races
				go func() {
					callbackStruct.callback(jsonArray...)
					if callbackStruct.oneTime {
						if verbose {
							log.Printf("Removing one-time callback \n")
						}
						deleteCallback(jsonArray[0])
					}
					if verbose {
						log.Printf("Current callbacks count %d \n", len(callbacks))
					}
				}()

			} else {
				log.Printf("Can't find callback for %s \n", jsonArray[0])
			}
		}

		document.ClientDisconnected = true

		for callbackID, callbackStruct := range callbacks {
			if callbackStruct.connection == nil {
				log.Fatalln("This should not happen")
			}
			if callbackStruct.connection == connection {
				if verbose {
					log.Printf("Removing callback %s for disconnected client\n", callbackID)
				}
				deleteCallback(callbackID)
			}
		}

		if verbose {
			log.Printf("Current callbacks count %d \n", len(callbacks))
		}
	}
}
