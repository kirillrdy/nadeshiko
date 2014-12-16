package html

import "bytes"

type Node struct {
	nodeType         string
	Attributes       map[string]string
	children         []Node
	text             string
	headTagMetaMagic string // This is for html doctype
}

func (node Node) writeChildrenToBuffer(buffer *bytes.Buffer) {
	for i := range node.children {
		node.children[i].writeToBuffer(buffer)
	}
}

func (node Node) attributesAsString() string {
	var result bytes.Buffer
	for attribute, value := range node.Attributes {
		result.WriteString(" ")
		result.WriteString(attribute)
		result.WriteString("=\"")
		result.WriteString(value)
		result.WriteString("\"")
	}
	return result.String()
}

func (node Node) String() string {
	buffer := &bytes.Buffer{}
	node.writeToBuffer(buffer)
	return buffer.String()
}

func (node Node) writeToBuffer(buffer *bytes.Buffer) {

	//TODO need to escape html
	buffer.WriteString("<")
	buffer.WriteString(node.headTagMetaMagic)
	buffer.WriteString(node.nodeType)
	buffer.WriteString(node.attributesAsString())
	buffer.WriteString(">")

	if len(node.children) > 0 {
		node.writeChildrenToBuffer(buffer)
	} else {
		buffer.WriteString(node.text)
	}

	buffer.WriteString("</")
	buffer.WriteString(node.nodeType)
	buffer.WriteString(">")

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
