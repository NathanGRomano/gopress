package gopress

import "container/list"

func Gopress() *Router {
	var router = new(Router)
	return router
}

type Router struct {
	fns *list.List
}

func (r *Router) Fns() *list.List {
	if r.fns == nil {
		r.fns = list.New()
	}
	return r.fns
}

func (r *Router) Use(fn func()) *Router {
	r.Fns().PushBack(fn)
	return r
}
