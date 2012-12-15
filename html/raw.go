package html

import "io/ioutil"

type Raw string

func (html Raw) Html() (data string) {
	data = string(html)
	return
}

func Template(name string) Raw {
	content, _ := ioutil.ReadFile("template/"+name)
	return Raw(string(content))
}
