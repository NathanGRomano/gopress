package gopress

type Router struct {
	caseSensitive bool
	strict        bool
	mergeParams   bool
	params        map[string][]ParamFn
	_params       []ParamFn
	stack         []Layer
}

func NewRouter(options map[string]bool) *Router {
	router := new(Router)
	router.caseSensitive = false
	router.strict = false
	router.mergeParams = false
	router.params = make(map[string][]ParamFn, 1)
	router._params = make([]ParamFn, 1)
	router.stack = make([]Layer, 1)
	if options != nil {
		if v, ok := options["caseSensitive"]; ok {
			router.caseSensitive = v
		}
		if v, ok := options["strict"]; ok {
			router.strict = v
		}
		if v, ok := options["mergeParams"]; ok {
			router.mergeParams = v
		}
	}
	return router
}
