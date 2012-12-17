package	nadeshiko

import "fmt"

func (element JQuerySelectedElements) Draggable() {
	string_to_send := fmt.Sprintf("$('%s').draggable()",element.selector)
	element.connection.SendMessage(string_to_send)
}
func (element JQuerySelectedElements) Sortable() {
	string_to_send := fmt.Sprintf("$('%s').sortable()",element.selector)
	element.connection.SendMessage(string_to_send)
}
