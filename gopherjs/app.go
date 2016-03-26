package gopherjs

import (
	"github.com/kirillrdy/nadeshiko/html"
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
}

func (app App) compiledJsFileName() string {
	return filepath.Base(app.PackageName) + ".js"
}

//CompiledJsFileWebPath is a path via which compile js app will be served
func (app App) CompiledJsFileWebPath() string {
	return "/" + app.compiledJsFileName()
}

func (app App) compile() error {
	now := time.Now()
	//TODO something that uses gopherjs as a library
	//TODO add -m (minified support)
	log.Printf("Compiling JS Application: %#v", app.PackageName)
	command := exec.Command("gopherjs", "build", app.PackageName)
	//TODO return these errors to the page  ( in dev mode )
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	log.Printf("Compilation took %v", time.Since(now))
	return err
}

//Handler returns handler that will serve a given gopherjs app
func (app App) Handler() http.HandlerFunc {

	return func(response http.ResponseWriter, request *http.Request) {
		//TODO dont do this here
		err := app.compile()
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
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
		http.ServeFile(response, request, app.compiledJsFileName())
	}
}
