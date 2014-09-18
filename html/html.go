package html

import "fmt"

type Html struct {
	Title string
	Head  []Node
	Body  []Node
}

func (html Html) bodyAsString() string {
	children_as_string := ""
	for _, child := range html.Body {
		children_as_string += child.String()
	}
	return children_as_string
}

func (html Html) String() string {
	body := `<!DOCTYPE html>
    <head>
        <meta charset="utf-8">
        <title>%s</title>
        <meta name="description" content="">
        <meta name="viewport" content="width=device-width, initial-scale=1">
    </head>
    <body>
		%s
    </body>
</html>
`
	return fmt.Sprintf(body, html.Title, html.bodyAsString())
}

func Span() Node {
	return Node{Type: "span"}
}
func P() Node {
	return Node{Type: "p"}
}
func Div() Node {
	return Node{Type: "div"}
}

func Body() Node {
	return Node{Type: "body"}
}

func Title() Node {
	return Node{Type: "title"}
}
