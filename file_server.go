package nadeshiko

import (
	"errors"
	"os"
	"path"
	"runtime"
)

//Turns on extra debugging
var verbose bool

func findStaticFile(file string) (string, error) {
	for _, dir := range publicDirs() {
		path := dir + file
		if stat, err := os.Stat(path); !os.IsNotExist(err) && !stat.IsDir() {
			return path, nil
		}

	}
	return "", errors.New("File not found")
}

func publicDirs() []string {
	return []string{"public", nadeshikoPublicDir()}
}

func nadeshikoPublicDir() string {
	_, current_file, _, _ := runtime.Caller(0)
	package_dir := path.Dir(current_file)
	return package_dir + "/public"
}
