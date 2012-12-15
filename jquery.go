package nadeshiko

import "fmt"
//import "math/rand"
import "strconv"
import "time"


type JQuerySelectedElements struct {
	selector string
	ws       *WebsocketConnection
}

func (ws *WebsocketConnection) JQuery(selector string) (element JQuerySelectedElements) {
	element.selector = selector
	element.ws = ws
	return
}

type HTMLAble interface {
	Html() string
}


//TODO DRY these 3 funcs Append,SetVal, SetText, Empty
func (element JQuerySelectedElements) Append(content HTMLAble) {
	string_content := strconv.Quote(content.Html())
	string_to_send := fmt.Sprintf("$('%s').append(%s)", element.selector, string_content)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) AppendString(content string) {
	string_content := strconv.Quote(content)
	string_to_send := fmt.Sprintf("$('%s').append(%s)", element.selector, string_content)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) SetVal(new_value string) {
	string_content := strconv.Quote(new_value)
	string_to_send := fmt.Sprintf("$('%s').val(%s)", element.selector, string_content)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) SetText(new_value string) {
	string_content := strconv.Quote(new_value)
	string_to_send := fmt.Sprintf("$('%s').text(%s)", element.selector, string_content)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) Empty() {
	string_to_send := fmt.Sprintf("$('%s').empty()", element.selector)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) Remove() {
	string_to_send := fmt.Sprintf("$('%s').remove()", element.selector)
	element.ws.SendMessage(string_to_send)
}



func (element JQuerySelectedElements) Click(callback func()) {
	now := time.Now()
	random_number := now.UnixNano()
	random_string := fmt.Sprintf("%x",random_number)

	Callbacks[random_string] = OverSocketCallback{element.ws, false, func(...string) {
		callback()
	}}

	string_to_send := fmt.Sprintf("$('%s').click(function(){ ws.send(JSON.stringify([\"%x\"])); });", element.selector, random_number)
	element.ws.SendMessage(string_to_send)
}

//TODO refactor function body, not DRY
func (element JQuerySelectedElements) Keydown(callback func(int)) {
	now := time.Now()
	//TODO get better way of generating uniq number
	random_number := now.UnixNano()
	random_string := fmt.Sprintf("%x",random_number)

	Callbacks[random_string] = OverSocketCallback{element.ws, false, func(vals ...string) {
		key, _ := strconv.Atoi(vals[1])
		callback(key)
	}}

	string_to_send := fmt.Sprintf("$('%s').keydown(function(event){ ws.send(JSON.stringify([\"%x\",event.which.toString()])); });", element.selector, random_number)
	element.ws.SendMessage(string_to_send)
}

func (element JQuerySelectedElements) GetVal(callback func(string)) {
	//TODO DRY this
	now := time.Now()
	random_number := now.UnixNano()
	random_string := fmt.Sprintf("%x",random_number)

	Callbacks[random_string] = OverSocketCallback{element.ws, true, func(vals ...string){
		callback(vals[1])
	}}

	string_to_send := fmt.Sprintf("ws.send( JSON.stringify([\"%x\",$('%s').val()])); ", random_number, element.selector)
	element.ws.SendMessage(string_to_send)
}
