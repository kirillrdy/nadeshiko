package html

import (
	"strings"

	"github.com/sparkymat/webdsl/css"
)

func (node Node) Src(value string) Node {
	return node.Attribute("src", value)
}

func (node Node) Style(value string) Node {
	return node.Attribute("style", value)
}

func (node Node) Id(value string) Node {
	return node.Attribute("id", value)
}

func (node Node) Class(values ...css.Class) Node {
	var valuesForJoin []string
	for _, value := range values {
		valuesForJoin = append(valuesForJoin, string(value))
	}
	attrValue := strings.Join(valuesForJoin, " ")
	return node.Attribute("class", attrValue)
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

func (node Node) Name(value string) Node {
	return node.Attribute("name", value)
}

func (node Node) For(value string) Node {
	return node.Attribute("for", value)
}

func (node Node) Method(value string) Node {
	return node.Attribute("method", value)
}

func (node Node) Selected(value string) Node {
	return node.Attribute("selected", value)
}

func (node Node) Align(value string) Node {
	return node.Attribute("align", value)
}

func (node Node) Action(value string) Node {
	return node.Attribute("action", value)
}

func (node Node) Placeholder(value string) Node {
	return node.Attribute("placeholder", value)
}
