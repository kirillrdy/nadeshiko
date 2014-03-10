package html

//I'm reconsidering if its better to deal with this in go way
//
//import (
//	"bytes"
//	"io/ioutil"
//	"html/template"
//	"runtime"
//	"path"
//)
//
//func LoadTemplate(name string) string{
//	_, caller_file_name, _, _ := runtime.Caller(1)
//	dir := path.Dir(caller_file_name)
//
//	path := path.Join(dir,"template/", name)
//	content, err := ioutil.ReadFile(path)
//	if err != nil { panic(err) }
//	return string(content)
//}
//
//func RenderString(content string, data interface{}) string {
//	//TODO be creative with name
//	tmpl, err := template.New("").Parse(content)
//	if err != nil { panic(err) }
//	var doc bytes.Buffer
//	err = tmpl.Execute(&doc, data)
//	if err != nil { panic(err) }
//	return doc.String()
//}
//
//func RenderTemplate(name string, data interface{}) string{
//	return RenderString(LoadTemplate(name),data)
//}
