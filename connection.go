package nadeshiko

import "code.google.com/p/go.net/websocket"
import "fmt"

type WebsocketConnection websocket.Conn

type Connection struct {
	websocket *WebsocketConnection
	currentActivity *Activity
}

func (connection *Connection) SendMessage(message string) {
	real_websocket := websocket.Conn(*connection.websocket)
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
	if Verbose {
		fmt.Printf("send: %s\n", message)
	}
}

