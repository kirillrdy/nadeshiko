package html

import "testing"

func TestIf(t *testing.T) {

	condition := true

	if generateDiv(condition) != "<div><span>Hello</span></div>" {
		t.Errorf("Failed to generate correct span got %#v", generateDiv(condition))
	}

	condition = false

	if generateDiv(condition) != "<div></div>" {
		t.Errorf("Failed to generate correct span got %#v", generateDiv(condition))
	}

}

func generateDiv(condition bool) string {
	return Div().Children(
		If(condition).Then(
			Span().Text("Hello"),
		).Nodes()...,
	).String()
}
