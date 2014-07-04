package gopress

func Gopress() *Router {
	var router = new(Router)
	return router
}

type Router struct {
	fns []func()
}

func (r *Router) Use(fn func()) *Router {
	r.fns = append(r.fns, fn)
	return r
}
