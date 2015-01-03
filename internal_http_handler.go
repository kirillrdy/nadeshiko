package nadeshiko

import (
	"log"
	"net/http"
	"time"
)

type internalHttpHandler struct {
	routes routes
}

func (h internalHttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	start_of_request := time.Now()

	response.Header().Set("Server", "Nadeshiko "+Version)

	file_path, err := findStaticFile(request.URL.Path)
	if err != nil {
		for _, route := range h.routes {

			//TODO also match request types
			if request.URL.Path == route.Path && request.Method == route.Method {
				log.Printf("%s: %q \n", request.Method, request.URL.Path)
				route.Handler(response, request)
				log.Printf("time taken: %s \n\n", time.Since(start_of_request).String())
				return
			}
		}
		//Last resort is 404
		http.NotFound(response, request)

	} else {
		http.ServeFile(response, request, file_path)
		return
	}

}
