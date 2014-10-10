package html

import (
	"bytes"
	"fmt"
)

type Node struct {
	nodeType         string
	Attributes       map[string]string
	children         []Node
	text             string
	headTagMetaMagic string // This is for html doctype
}

func (node Node) ChildrenAsString() string {
	var buffer bytes.Buffer
	for _, child := range node.children {
		buffer.WriteString(child.String())
	}
	return buffer.String()
}

func (node Node) AttributesAsString() string {
	var result bytes.Buffer
	for attribute, value := range node.Attributes {
		result.WriteString(fmt.Sprintf(" %s=\"%s\"", attribute, value))
	}
	return result.String()
}

func (node Node) String() string {
	var node_text string
	if len(node.children) > 0 {
		node_text = node.ChildrenAsString()
	} else {
		node_text = node.text
	}
	text := `<%s%s%s>%s</%s>`
	//TODO need to escape html
	return fmt.Sprintf(text, node.headTagMetaMagic, node.nodeType, node.AttributesAsString(), node_text, node.nodeType)
}

func (node Node) Text(text string) Node {
	node.text = text
	return node
}

func (node Node) Children(children ...Node) Node {
	node.children = append(node.children, children...)
	return node
}

func (node Node) Attribute(attribute, value string) Node {
	if node.Attributes == nil {
		node.Attributes = make(map[string]string)
	}
	node.Attributes[attribute] = value
	return node
}
