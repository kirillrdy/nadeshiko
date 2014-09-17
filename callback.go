package nadeshiko

import (
	"sync"

	"code.google.com/p/go.net/websocket"
)

type overSocketCallback struct {
	connection *websocket.Conn
	oneTime    bool
	callback   func(...string)
}

var callbacks = make(map[string]overSocketCallback)
var callbacksMutex sync.Mutex

func deleteCallback(callbackId string) {
	callbacksMutex.Lock()
	delete(callbacks, callbackId)
	callbacksMutex.Unlock()
}
