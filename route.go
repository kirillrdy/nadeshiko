package nadeshiko

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}
