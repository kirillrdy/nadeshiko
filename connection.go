package nadeshiko

import (
	"log"
	"runtime"
	"strings"
	"sync"

	"code.google.com/p/go.net/websocket"
)

type Connection struct {
	websocket           *websocket.Conn
	mutex               sync.Mutex
	current_transaction []string
	in_transaction      bool
}

func (connection *Connection) StartBuffer() {
	connection.in_transaction = true
}

func (connection *Connection) FlushBuffer() {
	connection.in_transaction = false
	if len(connection.current_transaction) != 0 {
		connection.SendMessage(strings.Join(connection.current_transaction, ";"))
		connection.current_transaction = []string{}
	}
}

func (connection *Connection) SendMessage(message string) {

	connection.mutex.Lock()
	if connection.in_transaction {
		connection.current_transaction = append(connection.current_transaction, message)
	} else {
		err := websocket.Message.Send(connection.websocket, message)

		if err != nil {
			log.Printf("runtime.Goexit '%s'\n", err)
			runtime.Goexit()

		}
	}

	connection.mutex.Unlock()

	// This is done to cleanup timers that are not terminated
	// but will try to send on close sockets
	// perhaps we can do this when we change activities
	//TODO handle errors other than send on closed connection

	if Verbose {
		log.Printf("send: %s\n", message)
	}
}
