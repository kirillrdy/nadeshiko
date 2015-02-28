package html

type ifWrapper struct {
	condition bool
	thenNodes []Node
	elseNodes []Node
}

func If(condition bool) ifWrapper {
	return ifWrapper{condition: condition}
}

func (wrapper ifWrapper) Then(nodes ...Node) ifWrapper {
	wrapper.thenNodes = nodes
	return wrapper
}

func (wrapper ifWrapper) Else(nodes ...Node) ifWrapper {
	wrapper.thenNodes = nodes
	return wrapper
}

func (wrapper ifWrapper) Nodes() []Node {
	if wrapper.condition {
		return wrapper.thenNodes
	} else {
		return wrapper.elseNodes
	}
}
