package gopress

type Request struct {
}

type Response struct {
}

type Cb func(err interface{})

type Fn func(req Request, res Response, next Cb)

type ErrFn func(err interface{}, req Request, res Resopnse, next Cb)

type Application struct {
	cache    map[string]interface{}
	settings map[string]interface{}
	engines  map[string]interface{}
}

func Gopress() *Application {
	var app = new(Application)
	app.Init()
	return app
}

func (a *Application) Init() {
	a.settings = make(map[string]interface{})
	a.cache = make(map[string]interface{})
	a.engines = make(map[string]interface{})
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

func (a *Application) SetEngine(name string, fn Fn) *Application {
	if !HasPrefix(name, ".") {
		name = "." + name
	}
	a.engines[name] = fn
	return a
}
