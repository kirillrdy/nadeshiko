package nadeshiko

import "fmt"

func (document *Document) SetLocalStorageItem(key, val string) {
	string_to_send := fmt.Sprintf("localStorage.setItem('%s','%s')", key, val)
	document.SendMessage(string_to_send)
}

func (document *Document) GetLocalStorageItem(key string, callback func(string)) {
	random_string := generateCallbackId()
	callbacks[random_string] = overSocketCallback{oneTime: true, callback: func(vals ...string) {
		callback(vals[1])
	}}

	string_to_send := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",localStorage.getItem('%s')])); ", random_string, key)
	document.SendMessage(string_to_send)
}
