package html

func (node Node) Src(value string) Node {
	return node.Attribute("src", value)
}

func (node Node) Style(value string) Node {
	return node.Attribute("style", value)
}

func (node Node) Id(value string) Node {
	return node.Attribute("id", value)
}

func (node Node) Class(value string) Node {
	return node.Attribute("class", value)
}

func (node Node) Type(value string) Node {
	return node.Attribute("type", value)
}

func (node Node) Value(value string) Node {
	return node.Attribute("value", value)
}

func (node Node) Href(value string) Node {
	return node.Attribute("href", value)
}

func (node Node) Media(value string) Node {
	return node.Attribute("media", value)
}

func (node Node) Rel(value string) Node {
	return node.Attribute("rel", value)
}
