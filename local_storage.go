package nadeshiko

import "fmt"

func (connection *Connection) SetLocalStorageItem(key, val string) {
	string_to_send := fmt.Sprintf("localStorage.setItem('%s','%s')", key, val)
	connection.SendMessage(string_to_send)
}

func (connection *Connection) GetLocalStorageItem(key string, callback func(string)) {
	random_string := generateCallbackId()
	callbacks[random_string] = overSocketCallback{connection, true, func(vals ...string) {
		callback(vals[1])
	}}

	string_to_send := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",localStorage.getItem('%s')])); ", random_string, key)
	connection.SendMessage(string_to_send)
}
