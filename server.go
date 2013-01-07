package nadeshiko

import (
	"log"
	"net/http"
)

const NADESHIKO_VERSION = "0.1.0"

type OverSocketCallback struct {
	connection  *Connection
	OneTimeOnly bool
	Callback    func(...string)
}

var Callbacks map[string]OverSocketCallback
var Verbose bool

func fileServer(w http.ResponseWriter, req *http.Request) {
	log.Printf("=======================================================\n")
	requested_path := req.RequestURI

	if requested_path == "/" {
		requested_path = "/base.html"
	}

	log.Printf("GET: %q \n", requested_path)
	log.Printf("User-Agent: %s \n", req.Header["User-Agent"])

	w.Header().Set("Server", "Nadeshiko "+NADESHIKO_VERSION)

	http.ServeFile(w, req, "public"+requested_path)

}


