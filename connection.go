package nadeshiko

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"runtime"
	"sync"
)

type WebsocketConnection websocket.Conn

type Connection struct {
	websocket       *WebsocketConnection
	mutex			sync.Mutex
}

func (connection *Connection) SendMessage(message string) {

	connection.mutex.Lock()
	real_websocket := websocket.Conn(*connection.websocket)
	err := websocket.Message.Send(&real_websocket, message)

	connection.mutex.Unlock()

	// This is done to cleanup timers that are not terminated
	// but will try to send on close sockets
	// perhaps we can do this when we change activities
	//TODO handle errors other than send on closed connection
	if err != nil {
		log.Printf("runtime.Goexit '%s'\n", err)
		runtime.Goexit()

	}

	if Verbose {
		log.Printf("send: %s\n", message)
	}
}
