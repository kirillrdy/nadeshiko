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

func Table() Node {
	return Node{Type: "table"}
}

func Thead() Node {
	return Node{Type: "thead"}
}

func Tbody() Node {
	return Node{Type: "tbody"}
}

func Th() Node {
	return Node{Type: "th"}
}

func Tr() Node {
	return Node{Type: "tr"}
}

func Td() Node {
	return Node{Type: "td"}
}

func Button() Node {
	return Node{Type: "button"}
}

func A() Node {
	return Node{Type: "a"}
}

func Input() Node {
	return Node{Type: "input"}
}

func Img() Node {
	return Node{Type: "img"}
}
