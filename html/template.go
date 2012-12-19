package html

import "bytes"
import "io/ioutil"
import "html/template"

func LoadTemplate(name string) string{
	content, err := ioutil.ReadFile("template/"+name)
	if err != nil { panic(err) }
	return string(content)
}

func RenderString(content string, data interface{}) string {
	//TODO be creative with name
	tmpl, err := template.New("").Parse(content)
	if err != nil { panic(err) }
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)
	if err != nil { panic(err) }
	return doc.String()
}

func RenderTemplate(name string, data interface{}) string{
	return RenderString(LoadTemplate(name),data)
}
