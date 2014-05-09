package nadeshiko

type Route struct {
	Path    string
	Method  string
	Handler HttpHander
}
