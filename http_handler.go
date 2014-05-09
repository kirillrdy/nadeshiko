package nadeshiko

import (
	"log"
	"net/http"
	"os"
	"time"
)

type HttpHander func(http.ResponseWriter, *http.Request)

type httpHandler struct {
	routes Routes
}

func (h httpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	start_of_request := time.Now()

	file_path := "public" + request.URL.Path
	if stat, err := os.Stat(file_path); os.IsNotExist(err) || stat.IsDir() {
		for _, route := range h.routes {

			//TODO also match request types
			if request.URL.Path == route.Path {
				log.Printf("%s request: %v \n", request.Method, request.URL.Path)
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
