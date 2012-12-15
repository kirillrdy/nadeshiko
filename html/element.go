package html

import "fmt"

type Element struct {
	NodeType  string
	Attributes Attributes
}

func (element Element) Html() string {
	return fmt.Sprintf("<%s %s></%s>", element.NodeType, element.Attributes, element.NodeType)
}

