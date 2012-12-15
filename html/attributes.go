package html

import "fmt"

type Attributes map[string]string

func (attributes Attributes) String() (body string) {

	for k, v := range attributes {
		attribute_part := fmt.Sprintf("%s=\"%s\"", k, v)
		body = body + attribute_part
	}
	return
}
