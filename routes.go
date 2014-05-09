package nadeshiko

const GET = "GET"
const POST = "POST"

type Routes []Route

func (routes *Routes) Get(path string, handler HttpHander) {
	route := Route{path, GET, handler}
	*routes = append(*routes, route)
}

func (routes *Routes) Post(path string, handler HttpHander) {
	route := Route{path, POST, handler}
	*routes = append(*routes, route)
}
