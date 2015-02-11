package nadeshiko

import "net/http"

type route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}
