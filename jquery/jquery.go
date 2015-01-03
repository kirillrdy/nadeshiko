package jquery

import (
	"net/http"
	"path"
	"runtime"
)

const jqueryFileName = "jquery-2.1.1.min.js"
const WebPath = "/" + jqueryFileName

func FileHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, filePath())
}

func filePath() string {
	_, current_file, _, _ := runtime.Caller(0)
	package_dir := path.Dir(current_file)
	return package_dir + WebPath
}
