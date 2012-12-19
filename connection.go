package nadeshiko

import "code.google.com/p/go.net/websocket"
import "fmt"
import "runtime"

type WebsocketConnection websocket.Conn

type Connection struct {
	websocket *WebsocketConnection
	currentActivity *Activity
}

func (connection *Connection) SendMessage(message string) {
	real_websocket := websocket.Conn(*connection.websocket)
	err := websocket.Message.Send(&real_websocket, message)

	// This is done to cleanup timers that are not terminated
	// but will try to send on close sockets
	// perhaps we can do this when we change activities
	if err != nil { runtime.Goexit() }

	if Verbose {
		fmt.Printf("send: %s\n", message)
	}
}

