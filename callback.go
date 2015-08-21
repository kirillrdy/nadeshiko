package nadeshiko

import (
	"sync"

	"golang.org/x/net/websocket"
)

type overSocketCallback struct {
	connection *websocket.Conn
	oneTime    bool
	callback   func(...string)
}

var callbacks = make(map[string]overSocketCallback)
var callbacksMutex sync.Mutex

func deleteCallback(callbackID string) {
	callbacksMutex.Lock()
	delete(callbacks, callbackID)
	callbacksMutex.Unlock()
}
