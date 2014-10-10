package html

import "testing"

func BenchmarkNodeToString(b *testing.B) {
	node := Div()
	for i := 0; i < 1000; i++ {
		nested := Div()
		for i := 0; i < 1000; i++ {
			nested.Children(Span().Text("Some text"))
		}

		node.Children(nested)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node.String()
	}
}
