package nadeshiko

import (
	"log"
	"net/http"
	"time"
)

type internalHTTPHandler struct {
}

//Going to get rid of internal router soon anyways
func (h internalHTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	startOfRequest := time.Now()

	response.Header().Set("Server", "Nadeshiko "+Version)

	filePath, err := findStaticFile(request.URL.Path)
	if err != nil {
		for _, route := range defaultRoutes {

			if request.URL.Path == route.Path && request.Method == route.Method {
				log.Printf("%s: %q \n", request.Method, request.URL.Path)
				route.Handler(response, request)
				log.Printf("time taken: %s \n\n", time.Since(startOfRequest).String())
				return
			}
		}
		//Last resort is 404
		http.NotFound(response, request)

	} else {
		http.ServeFile(response, request, filePath)
		return
	}

}
