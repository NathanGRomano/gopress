package gopress

type Route struct {
	path       string
	stack      []Layer
	methods    map[string]*Fn
	errMethods map[string]*ErrFn
}

func NewRoute(path string) *Route {
	var route = new(Route)
	if len(path) == 0 {
		route.path = "/"
	} else {
		route.path = path
	}
	route.methods = make(map[string]*Fn)
	route.errMethods = make(map[string]*ErrFn)
	return route
}
