package html

import (
	"fmt"
	"strings"

	"github.com/sparkymat/webdsl/css"
)

func (node Node) Src(value string) Node {
	return node.Attribute("src", value)
}

func (node Node) Style(value string) Node {
	return node.Attribute("style", value)
}

func (node Node) Id(id css.Id) Node {
	return node.Attribute("id", string(id))
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

func (node Node) Width(value uint) Node {
	return node.Attribute("width", fmt.Sprintf("%v", value))
}

func (node Node) Height(value uint) Node {
	return node.Attribute("height", fmt.Sprintf("%v", value))
}

func (node Node) WidthFloat(value float64) Node {
	return node.Attribute("width", fmt.Sprintf("%v", value))
}

func (node Node) HeightFloat(value float64) Node {
	return node.Attribute("height", fmt.Sprintf("%v", value))
}

func (node Node) Title(value string) Node {
	return node.Attribute("title", value)
}

func (node Node) Controls() Node {
	return node.Attribute("controls", "")
}

func (node Node) Autoplay() Node {
	return node.Attribute("autoplay", "")
}

func (node Node) Label(value string) Node {
	return node.Attribute("label", value)
}

func (node Node) Kind(value string) Node {
	return node.Attribute("kind", value)
}

func (node Node) Srclang(value string) Node {
	return node.Attribute("srclang", value)
}

func (node Node) Default() Node {
	return node.Attribute("default", "")
}

func (node Node) Multiple() Node {
	return node.Attribute("multiple", "")
}

func (node Node) Accept(value string) Node {
	return node.Attribute("accept", value)
}

func (node Node) D(value string) Node {
	return node.Attribute("d", value)
}

func (node Node) Fill(value string) Node {
	return node.Attribute("fill", value)
}

func (node Node) Stroke(value string) Node {
	return node.Attribute("stroke", value)
}
