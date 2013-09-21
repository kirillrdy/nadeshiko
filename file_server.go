package nadeshiko

import (
	"log"
	"net/http"
)


var Verbose bool

func fileServer(w http.ResponseWriter, req *http.Request) {
	log.Printf("=======================================================\n")
	requested_path := req.RequestURI

	if requested_path == "/" {
		requested_path = "/index.html"
	}

	log.Printf("GET: %q \n", requested_path)
	log.Printf("User-Agent: %s \n\n", req.Header["User-Agent"])

	w.Header().Set("Server", "Nadeshiko " + NADESHIKO_VERSION)

	//TODO check nadeshiko bundle public and then serve from local public is availible
	http.ServeFile(w, req, "nadeshiko/public" + requested_path)

}
