package nadeshiko

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

type JQuerySelector struct {
	selector string
	document *Document
}

func (document *Document) JQuery(selector css.Selector) (element JQuerySelector) {
	element.selector = selector.Selector()
	element.document = document
	return
}

//func (element JQuerySelector) Write(content []byte) (int, error) {
//	element.Append(string(content))
//	return len(content), nil
//}

// Html
func (element JQuerySelector) Html(content html.Node) {
	element.oneArgumentMethod("html", content.String())
}

func (element JQuerySelector) Append(content html.Node) {
	element.oneArgumentMethod("append", content.String())
}

func (element JQuerySelector) Before(content html.Node) {
	element.oneArgumentMethod("before", content.String())
}

func (element JQuerySelector) Prepend(content html.Node) {
	element.oneArgumentMethod("prepend", content.String())
}

func (element JQuerySelector) PrependString(content string) {
	element.oneArgumentMethod("prepend", content)
}

func (element JQuerySelector) SetVal(new_value string) {
	element.oneArgumentMethod("val", new_value)
}

func (element JQuerySelector) SetCss(attr, new_value string) {
	element.twoArgumentMethod("css", attr, new_value)
}

func (element JQuerySelector) SetAttr(attr, new_value string) {
	element.twoArgumentMethod("attr", attr, new_value)
}

func (element JQuerySelector) SetText(new_value string) {
	element.oneArgumentMethod("text", new_value)
}

func (element JQuerySelector) Empty() {
	element.zeroArgumentMethod("empty")
}

func (element JQuerySelector) Remove() {
	element.zeroArgumentMethod("remove")
}

func (element JQuerySelector) Click(callback func()) {
	element.zeroArgumentMethodWithCallback("click", callback)
}

func (element JQuerySelector) On(eventName string, callback func()) {
	element.oneArgumentMethodWithCallback("on", eventName, callback)
}

func (element JQuerySelector) Change(callback func()) {
	element.zeroArgumentMethodWithCallback("change", callback)
}

//TODO refactor function body, not DRY
func (element JQuerySelector) Keypress(callback func(int)) {

	callback_id := generateCallbackId()

	callbacks[callback_id] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	string_to_send := fmt.Sprintf("$('%s').keypress(function(event){ ws.send(JSON.stringify([\"%s\",event.charCode.toString()])); });", element.selector, callback_id)
	element.document.SendMessage(string_to_send)
}

//TODO refactor function body, not DRY
func (element JQuerySelector) Keydown(callback func(int)) {

	callback_id := generateCallbackId()

	callbacks[callback_id] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	string_to_send := fmt.Sprintf("$('%s').keydown(function(event){ ws.send(JSON.stringify([\"%s\",event.keyCode.toString()])); });", element.selector, callback_id)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelector) GetVal(callback func(string)) {
	random_string := generateCallbackId()

	callbacks[random_string] = overSocketCallback{connection: element.document.websocket, oneTime: true, callback: func(vals ...string) {
		callback(vals[1])
	}}

	string_to_send := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",$('%s').val()])); ", random_string, element.selector)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelector) GetValChan() chan string {
	random_string := generateCallbackId()

	result := make(chan string)

	callback := func(vals ...string) {
		result <- vals[1]
	}

	callbacks[random_string] = overSocketCallback{connection: element.document.websocket, oneTime: true, callback: callback}

	string_to_send := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",$('%s').val()])); ", random_string, element.selector)
	element.document.SendMessage(string_to_send)
	return result
}

////////////////////////////////////////////////////////////////////////////////////
// Unexported functions go here

func generateCallbackId() string {
	now := time.Now()
	//TODO get better way of generating uniq number
	random_number := now.UnixNano()
	return fmt.Sprintf("%x", random_number)
}

func (element JQuerySelector) twoArgumentMethod(name, param1, param2 string) {
	quoted1 := strconv.Quote(param1)
	quoted2 := strconv.Quote(param2)
	string_to_send := fmt.Sprintf("$('%s').%s(%s,%s)", element.selector, name, quoted1, quoted2)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelector) oneArgumentMethod(name string, param string) {
	string_content := strconv.Quote(param)
	string_to_send := fmt.Sprintf("$('%s').%s(%s)", element.selector, name, string_content)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelector) zeroArgumentMethod(name string) {
	string_to_send := fmt.Sprintf("$('%s').%s()", element.selector, name)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelector) oneArgumentMethodWithCallback(name string, arg1 string, callback func()) {
	callbackID := generateCallbackId()

	callbacks[callbackID] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(...string) {
		callback()
	}}

	stringToSend := fmt.Sprintf("$('%s').%s('%s', function(){ ws.send(JSON.stringify([\"%s\"])); });", element.selector, name, arg1, callbackID)
	element.document.SendMessage(stringToSend)
}

func (element JQuerySelector) zeroArgumentMethodWithCallback(name string, callback func()) {
	callbackID := generateCallbackId()

	callbacks[callbackID] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(...string) {
		callback()
	}}

	stringToSend := fmt.Sprintf("$('%s').%s(function(){ ws.send(JSON.stringify([\"%s\"])); });", element.selector, name, callbackID)
	element.document.SendMessage(stringToSend)
}
