package nadeshiko

import (
	"log"
	"net/http"
	"os"
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

	log.Printf("GET: %q \n", requested_path)
	if Verbose {
		log.Printf("User-Agent: %s \n\n", req.Header["User-Agent"])
	}

	w.Header().Set("Server", "Nadeshiko "+NADESHIKO_VERSION)

	file_path := findStaticFile(req.URL.Path)
	http.ServeFile(w, req, file_path)

}

func findStaticFile(file string) string {
	for _, dir := range publicDirs() {
		path := dir + file
		if stat, err := os.Stat(path); !os.IsNotExist(err) && !stat.IsDir() {
			return path
		}

	}
	return ""
}

func publicDirs() []string {
	return []string{"public", nadeshikoPublicDir()}
}

func nadeshikoPublicDir() string {

	//TODO check nadeshiko bundle public and then serve from local public is availible
	_, current_file, _, _ := runtime.Caller(0)
	package_dir := path.Dir(current_file)
	return package_dir + "/public"
}
