package html

func Html() Node {
	//TODO Instead of doing he here we can just hard code assumption for html node in Serialiser
	return Node{nodeType: "html", headTagMetaMagic: "<!DOCTYPE html>"}
}

func Meta() Node {
	return Node{nodeType: "meta"}
}

func Span() Node {
	return Node{nodeType: "span"}
}
func P() Node {
	return Node{nodeType: "p"}
}
func Div() Node {
	return Node{nodeType: "div"}
}

func Nav() Node {
	return Node{nodeType: "nav"}
}

func Head() Node {
	return Node{nodeType: "head"}
}

func Body() Node {
	return Node{nodeType: "body"}
}

func Style() Node {
	return Node{nodeType: "style"}
}

func Script() Node {
	return Node{nodeType: "script"}
}

func Title() Node {
	return Node{nodeType: "title"}
}

func H1() Node {
	return Node{nodeType: "h1"}
}

func H2() Node {
	return Node{nodeType: "h2"}
}

func H3() Node {
	return Node{nodeType: "h3"}
}

func Svg() Node {
	return Node{nodeType: "svg"}
}

func Path() Node {
	return Node{nodeType: "path"}
}

func Table() Node {
	return Node{nodeType: "table"}
}

func Thead() Node {
	return Node{nodeType: "thead"}
}

func Tbody() Node {
	return Node{nodeType: "tbody"}
}

func Th() Node {
	return Node{nodeType: "th"}
}

func Tr() Node {
	return Node{nodeType: "tr"}
}

func Td() Node {
	return Node{nodeType: "td"}
}

func Button() Node {
	return Node{nodeType: "button"}
}

func A() Node {
	return Node{nodeType: "a"}
}

func Input() Node {
	return Node{nodeType: "input"}
}

func Img() Node {
	return Node{nodeType: "img"}
}

func I() Node {
	return Node{nodeType: "i"}
}

func Link() Node {
	return Node{nodeType: "link"}
}

func Ul() Node {
	return Node{nodeType: "ul"}
}

func Ol() Node {
	return Node{nodeType: "ol"}
}

func Li() Node {
	return Node{nodeType: "li"}
}

func Hr() Node {
	return Node{nodeType: "hr"}
}

func Form() Node {
	return Node{nodeType: "form"}
}

func Label() Node {
	return Node{nodeType: "label"}
}

func Pre() Node {
	return Node{nodeType: "pre"}
}

func Select() Node {
	return Node{nodeType: "select"}
}

func Option() Node {
	return Node{nodeType: "option"}
}

func Textarea() Node {
	return Node{nodeType: "textarea"}
}

func Video() Node {
	return Node{nodeType: "video"}
}

func Source() Node {
	return Node{nodeType: "source"}
}

func Track() Node {
	return Node{nodeType: "track"}
}

func Footer() Node {
	return Node{nodeType: "footer"}
}

func H4() Node {
	return Node{nodeType: "h4"}
}

func H5() Node {
	return Node{nodeType: "h5"}
}
