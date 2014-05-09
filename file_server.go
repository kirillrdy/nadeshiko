package nadeshiko

import (
	"log"
	"net/http"
	"path"
	"runtime"
)

var Verbose bool

func fileServer(w http.ResponseWriter, req *http.Request) {
	requested_path := req.RequestURI

	if Verbose {
		log.Printf("=======================================================\n")
		log.Println("Serving static content: " + requested_path)
	}

	if requested_path == "/" {
		requested_path = "/index.html"
	}

	log.Printf("HTTP GET: %q \n", requested_path)
	if Verbose {
		log.Printf("User-Agent: %s \n\n", req.Header["User-Agent"])
	}

	w.Header().Set("Server", "Nadeshiko "+NADESHIKO_VERSION)

	http.ServeFile(w, req, nadeshikoPublicDir()+requested_path)

}

func nadeshikoPublicDir() string {

	//TODO check nadeshiko bundle public and then serve from local public is availible
	_, current_file, _, _ := runtime.Caller(0)
	package_dir := path.Dir(current_file)
	return package_dir + "/public"
}
