package nadeshiko

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

type JQuerySelectedElements struct {
	selector string
	document *Document
}

func (document *Document) JQuery(selector css.Selector) (element JQuerySelectedElements) {
	element.selector = selector.Selector()
	element.document = document
	return
}

//func (element JQuerySelectedElements) Write(content []byte) (int, error) {
//	element.Append(string(content))
//	return len(content), nil
//}

func (element JQuerySelectedElements) Html(content html.Node) {
	element.oneArgumentMethod("html", content.String())
}

func (element JQuerySelectedElements) Append(content html.Node) {
	element.oneArgumentMethod("append", content.String())
}

//TODO get rid of this method, and figure out more neat way of chaining jquery methods
func (element JQuerySelectedElements) PrevRemove() {
	string_to_send := fmt.Sprintf("$('%s').prev().remove()", element.selector)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) Before(content string) {
	element.oneArgumentMethod("before", content)
}

func (element JQuerySelectedElements) Prepend(content string) {
	element.oneArgumentMethod("prepend", content)
}

func (element JQuerySelectedElements) PrependString(content string) {
	element.oneArgumentMethod("prepend", content)
}

func (element JQuerySelectedElements) SetVal(new_value string) {
	element.oneArgumentMethod("val", new_value)
}

func (element JQuerySelectedElements) SetCss(attr, new_value string) {
	element.twoArgumentMethod("css", attr, new_value)
}

func (element JQuerySelectedElements) SetAttr(attr, new_value string) {
	element.twoArgumentMethod("attr", attr, new_value)
}

func (element JQuerySelectedElements) SetText(new_value string) {
	element.oneArgumentMethod("text", new_value)
}

func (element JQuerySelectedElements) Empty() {
	element.zeroArgumentMethod("empty")
}

func (element JQuerySelectedElements) Remove() {
	element.zeroArgumentMethod("remove")
}

func (element JQuerySelectedElements) Click(callback func()) {
	element.zeroArgumentMethodWithCallback("click", callback)
}

func (element JQuerySelectedElements) On(eventName string, callback func()) {
	element.oneArgumentMethodWithCallback("on", eventName, callback)
}

func (element JQuerySelectedElements) Change(callback func()) {
	element.zeroArgumentMethodWithCallback("change", callback)
}

//TODO refactor function body, not DRY
func (element JQuerySelectedElements) Keypress(callback func(int)) {

	callback_id := generateCallbackId()

	callbacks[callback_id] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	string_to_send := fmt.Sprintf("$('%s').keypress(function(event){ ws.send(JSON.stringify([\"%s\",event.charCode.toString()])); });", element.selector, callback_id)
	element.document.SendMessage(string_to_send)
}

//TODO refactor function body, not DRY
func (element JQuerySelectedElements) Keydown(callback func(int)) {

	callback_id := generateCallbackId()

	callbacks[callback_id] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	string_to_send := fmt.Sprintf("$('%s').keydown(function(event){ ws.send(JSON.stringify([\"%s\",event.keyCode.toString()])); });", element.selector, callback_id)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) GetVal(callback func(string)) {
	random_string := generateCallbackId()

	callbacks[random_string] = overSocketCallback{connection: element.document.websocket, oneTime: true, callback: func(vals ...string) {
		callback(vals[1])
	}}

	string_to_send := fmt.Sprintf("ws.send( JSON.stringify([\"%s\",$('%s').val()])); ", random_string, element.selector)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) GetValChan() chan string {
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

func (element JQuerySelectedElements) twoArgumentMethod(name, param1, param2 string) {
	quoted1 := strconv.Quote(param1)
	quoted2 := strconv.Quote(param2)
	string_to_send := fmt.Sprintf("$('%s').%s(%s,%s)", element.selector, name, quoted1, quoted2)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) oneArgumentMethod(name string, param string) {
	string_content := strconv.Quote(param)
	string_to_send := fmt.Sprintf("$('%s').%s(%s)", element.selector, name, string_content)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) zeroArgumentMethod(name string) {
	string_to_send := fmt.Sprintf("$('%s').%s()", element.selector, name)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) oneArgumentMethodWithCallback(name string, arg1 string, callback func()) {
	callback_id := generateCallbackId()

	callbacks[callback_id] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(...string) {
		callback()
	}}

	string_to_send := fmt.Sprintf("$('%s').%s('%s', function(){ ws.send(JSON.stringify([\"%s\"])); });", element.selector, name, arg1, callback_id)
	element.document.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) zeroArgumentMethodWithCallback(name string, callback func()) {
	callback_id := generateCallbackId()

	callbacks[callback_id] = overSocketCallback{connection: element.document.websocket, oneTime: false, callback: func(...string) {
		callback()
	}}

	string_to_send := fmt.Sprintf("$('%s').%s(function(){ ws.send(JSON.stringify([\"%s\"])); });", element.selector, name, callback_id)
	element.document.SendMessage(string_to_send)
}
