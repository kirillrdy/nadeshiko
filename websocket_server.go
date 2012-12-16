package nadeshiko

import "code.google.com/p/go.net/websocket"
import "strconv"
import "encoding/json"
import "os"
import "fmt"

func websocketServer(ws *websocket.Conn) {
	fmt.Printf("New client connection on %#v\n", ws)

	connection := WebsocketConnection(*ws)

	//TODO check for DefaultActivity being set
	if DefaultActivity == nil {
		panic("Need to set DefaultActivity")
		os.Exit(1)
	}
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
			fmt.Printf("ERROR reading from socket: %v \n",err)
			break
		}

		if verbose {
			quoted_contet := strconv.Quote(buf)
			fmt.Printf("received: %s\n", quoted_contet)
		}

		var json_array []string
		json.Unmarshal([]byte(buf),&json_array)

		//This if statment not really needed,
		// since websocket.Message.Receive should catch most of errors
		// but good to keep for debugging
		if callbackStruct, ok := Callbacks[json_array[0]]; ok {
			callbackStruct.Callback(json_array...)
			if callbackStruct.OneTimeOnly {
				fmt.Printf("Removing one time %d \n",len(Callbacks))
				delete(Callbacks, json_array[0])
			}

			fmt.Printf("Current callbacks count %d \n",len(Callbacks))

		} else {
			fmt.Printf("Cant find callback for %s \n",json_array[0])
		}
	}

	fmt.Println("Client disconnected")

	//TODO find more efficient way of dealing with lots of notifications
	//TODO FIX FIX
	//for k, v := range Notifications {
	//	var new_list []WebsocketConnection
	//	for _, jFunction := range v {
	//		if jFunction("").ws != &wrapper_socket {
	//			new_list = append(new_list,jFunction)
	//		} else {
	//			fmt.Printf("Removing Notification '%s' for disconnected client %v\n", k, wrapper_socket)
	//		}
	//	}
	//	Notifications[k] = new_list
	//}

	for callback_id, callbackStruct := range Callbacks {
		if callbackStruct.ws == &connection {
			delete(Callbacks, callback_id)
			fmt.Printf("Removing callback %s for disconnected client\n",callback_id)
		}
	}
	fmt.Printf("Current callbacks count %d \n",len(Callbacks))


}
