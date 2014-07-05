package gopress

type Router struct {
	caseSensitive bool
	strict        bool
}

func NewRouter(options map[string]bool) *Router {
	router := new(Router)
	router.caseSensitive = false
	router.strict = false
	if options != nil {
		if v, ok := options["caseSensitive"]; ok {
			router.caseSensitive = v
		}
		if v, ok := options["strict"]; ok {
			router.strict = v
		}
	}
	return router
}
