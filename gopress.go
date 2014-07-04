package gopress

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
