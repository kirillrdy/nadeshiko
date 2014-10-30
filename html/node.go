package html

import (
	"bytes"
	"io"
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
	for i := range node.children {
		buffer.WriteString(node.children[i].String())
	}
	return buffer.String()
}

func (node Node) WriteChildrenToBuffer(buffer *bytes.Buffer) {
	for i := range node.children {
		node.children[i].WriteToBuffer(buffer)
	}
}

func (node Node) WriteChildren(writer io.Writer) {
	for i := range node.children {
		node.children[i].Write(writer)
	}
}

func (node Node) AttributesAsString() string {
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
	node.WriteToBuffer(buffer)
	return buffer.String()
}

func (node Node) Write(writer io.Writer) {

	//TODO need to escape html
	io.WriteString(writer, "<")
	io.WriteString(writer, node.headTagMetaMagic)
	io.WriteString(writer, node.nodeType)
	io.WriteString(writer, node.AttributesAsString())
	io.WriteString(writer, ">")

	if len(node.children) > 0 {
		node.WriteChildren(writer)
	} else {
		io.WriteString(writer, node.text)
	}

	io.WriteString(writer, "</")
	io.WriteString(writer, node.nodeType)
	io.WriteString(writer, ">")

}
func (node Node) WriteToBuffer(buffer *bytes.Buffer) {

	//TODO need to escape html
	buffer.WriteString("<")
	buffer.WriteString(node.headTagMetaMagic)
	buffer.WriteString(node.nodeType)
	buffer.WriteString(node.AttributesAsString())
	buffer.WriteString(">")

	if len(node.children) > 0 {
		node.WriteChildrenToBuffer(buffer)
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
