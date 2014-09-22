package html

func Html() Node {
	return Node{Type: "html", headTagMetaMagic: "!DOCTYPE "}
}

func Meta() Node {
	return Node{Type: "meta"}
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

func Head() Node {
	return Node{Type: "head"}
}

func Body() Node {
	return Node{Type: "body"}
}

func Style() Node {
	return Node{Type: "style"}
}

func Script() Node {
	return Node{Type: "script"}
}

func Title() Node {
	return Node{Type: "title"}
}

func H1() Node {
	return Node{Type: "h1"}
}

func H2() Node {
	return Node{Type: "h2"}
}

func H3() Node {
	return Node{Type: "h3"}
}

func Svg() Node {
	return Node{Type: "svg"}
}

func Path() Node {
	return Node{Type: "path"}
}
