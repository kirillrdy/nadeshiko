package gopherjs

import (
	"github.com/kirillrdy/web/html"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

//App is representation of Gopherjs Application
type App struct {
	PackageName string
}

//Mount registers GopherJS app using http.HandleFunc
func (app App) Mount(path string) {
	http.HandleFunc(path, app.Handler())
	http.HandleFunc(app.CompiledJsFileWebPath(), app.JsFileHandler())
	http.HandleFunc(app.CompiledJsMapFileWebPath(), app.JsMapFileHandler())
}

func (app App) compiledJsFileName() string {
	return filepath.Base(app.PackageName) + ".js"
}

func (app App) compiledJsMapFileName() string {
	return app.compiledJsFileName() + ".map"
}

//CompiledJsFileWebPath is a path via which compiled js app will be served
func (app App) CompiledJsFileWebPath() string {
	return "/" + app.compiledJsFileName()
}

//CompiledJsMapFileWebPath is a path via which compiled js map app will be served
func (app App) CompiledJsMapFileWebPath() string {
	return "/" + app.compiledJsMapFileName()
}

func (app App) compile(errorWriter io.Writer) error {
	now := time.Now()
	//TODO something that uses gopherjs as a library
	//TODO add -m (minified support)
	log.Printf("Compiling JS Application: %#v", app.PackageName)
	command := exec.Command("gopherjs", "build", app.PackageName)
	//TODO return these errors to the page  ( in dev mode )
	command.Stdout = errorWriter
	command.Stderr = errorWriter
	err := command.Run()
	log.Printf("Compilation took %v", time.Since(now))
	return err
}

//Handler returns handler that will serve a given gopherjs app
func (app App) Handler() http.HandlerFunc {

	return func(response http.ResponseWriter, request *http.Request) {

		//TODO this is clearly for dev only
		if request.URL.Path != "/" {
			log.Println(request.URL.Path)
			goPath := os.Getenv("GOPATH")
			http.ServeFile(response, request, goPath+"/src"+request.URL.Path)
			return
		}

		//TODO dont do this here
		err := app.compile(response)
		if err != nil {
			//http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		page := html.Html().Children(
			html.Head().Children(
				html.Script().Src(app.CompiledJsFileWebPath()),
			),
			html.Body(),
		)
		page.WriteTo(response)
	}
}

//JsFileHandler serves compiled js file to be served via CompiledJsFileWebPath()
func (app App) JsFileHandler() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Println("Serving js file")
		http.ServeFile(response, request, app.compiledJsFileName())
	}
}

//JsMapFileHandler serves compiled js file to be served via CompiledJsFileWebPath()
func (app App) JsMapFileHandler() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Println("Serving js map file")
		http.ServeFile(response, request, app.compiledJsMapFileName())
	}
}
