package gopress

type Router struct {
	options map[string]interface{}
}

func NewRouter(options *map[string]interface{}) *Router {
	router = new(Router)
	router.options = make(map[string]interface{})
	router.options["caseSensitive"] = false
	router.options["strict"] = false
	if options != nil {
		router.options["caseSensitive"] = options["caseSensitive"] || false
		router.options["strict"] = options["strict"] || false
	}
	return router
}
