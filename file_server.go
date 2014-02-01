package nadeshiko

import (
	"log"
	"net/http"
	"runtime"
	"path"
)


var Verbose bool

func fileServer(w http.ResponseWriter, req *http.Request) {
	if Verbose {
		log.Printf("=======================================================\n")
	}

	requested_path := req.RequestURI

	if requested_path == "/" {
		requested_path = "/index.html"
	}

	log.Printf("HTTP GET: %q \n", requested_path)
	if Verbose {
		log.Printf("User-Agent: %s \n\n", req.Header["User-Agent"])
	}

	w.Header().Set("Server", "Nadeshiko " + NADESHIKO_VERSION)

	//TODO check nadeshiko bundle public and then serve from local public is availible
	_, current_file, _, _ := runtime.Caller(0)
	package_dir := path.Dir(current_file)
	http.ServeFile(w, req, package_dir +"/public" + requested_path)

}
