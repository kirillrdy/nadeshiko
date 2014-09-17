package nadeshiko

import (
	"log"
	"net/http"
	"time"
)

type internalHttpHandler struct {
	routes routes
}

func (h internalHttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	start_of_request := time.Now()

	file_path := findStaticFile(request.URL.Path)
	if file_path == "" {
		for _, route := range h.routes {

			//TODO also match request types
			if request.URL.Path == route.Path {
				log.Printf("%s: %q \n", request.Method, request.URL.Path)
				route.Handler(writer, request)
				log.Printf("time taken: %s \n\n", time.Since(start_of_request).String())
				return
			}
		}
		//Last resort is 404
		http.NotFound(writer, request)

	} else {

		fileServer(writer, request)
		return
	}

}
