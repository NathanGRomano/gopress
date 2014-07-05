package gopress

import "strings"

type Fn func(req Request, res Response, next Cb)
type ErrFn func(err interface{}, req Request, res Response, next Cb)
type Cb func(err interface{})

type Application struct {
	cache    map[string]interface{}
	settings map[string]interface{}
	engines  map[string]*Fn
	router   *Router
}

func NewApplication() *Application {
	var app = new(Application)
	app.Init()
	return app
}

func (a *Application) Init() {
	a.settings = make(map[string]interface{})
	a.cache = make(map[string]interface{})
	a.engines = make(map[string]*Fn)
	a.DefaultConfiguration()
}

func (a *Application) DefaultConfiguration() {
}

func (a *Application) Set(name string, val interface{}) *Application {
	a.settings[name] = val
	return a
}

func (a *Application) Get(name string) interface{} {
	return a.settings[name]
}

func (a *Application) Enable(name string) *Application {
	return a.Set(name, true)
}

func (a *Application) Enabled(name string) bool {
	return a.Get(name) == true
}

func (a *Application) Disable(name string) *Application {
	return a.Set(name, false)
}

func (a *Application) Disabled(name string) bool {
	return a.Get(name) == false
}

func (a *Application) SetEngine(name string, fn *Fn) *Application {
	a.engines[prefix(name)] = fn
	return a
}

func (a *Application) Engine(name string) *Fn {
	return a.engines[prefix(name)]
}

func (a *Application) Router() *Router {
	if a.router == nil {
		conf := map[string]bool{
			"caseSensitive": a.Enabled("case sensitive routing"),
			"strict":        a.Enabled("strict routing"),
		}
		a.router = NewRouter(conf)
		// TODO we need to put in query and middleware
	}
	return a.router
}

func prefix(name string) string {
	if !strings.HasPrefix(name, ".") {
		name = "." + name
	}
	return name
}
