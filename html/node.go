package html

import "fmt"

type Node struct {
	nodeType         string
	Attributes       map[string]string
	children         []Node
	text             string
	headTagMetaMagic string // This is for html doctype
}

func (node Node) ChildrenAsString() string {
	children_as_string := ""
	for _, child := range node.children {
		children_as_string += child.String()
	}
	return children_as_string
}

func (node Node) AttributesAsString() string {
	result := ""
	for attribute, value := range node.Attributes {
		result += fmt.Sprintf(" %s=\"%s\"", attribute, value)
	}
	return result
}

func (node Node) String() string {
	node_text := ""
	if node.text == "" {
		node_text = node.ChildrenAsString()
	} else {
		node_text = node.text
	}
	text := `<%s%s%s>%s</%s>`
	return fmt.Sprintf(text, node.headTagMetaMagic, node.Type, node.AttributesAsString(), node_text, node.Type)
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
