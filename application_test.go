package gopress

import "testing"

func TestApplicationNewApplication(t *testing.T) {
	var app = NewApplication()
	if app == nil {
		t.Error("App should be defined")
	}
}

func TestApplicationInit(t *testing.T) {
	var app = new(Application)
	app.Init()
	if app.settings == nil {
		t.Error("settings are null")
	}
	if app.cache == nil {
		t.Error("cache is null")
	}
	if app.engines == nil {
		t.Error("engines are null")
	}
}

func TestApplicationDefaultConfiguration(t *testing.T) {
	var app = new(Application)
	app.Init()
	app.DefaultConfiguration()
}

func TestApplicationSet(t *testing.T) {
	app := NewApplication()
	key := "key"
	val := "value"
	app.Set(key, val)
	if app.settings[key] != val {
		t.Error("can't set value")
	}
}

func TestApplicationGet(t *testing.T) {
	app := NewApplication()
	key := "key"
	val := "value"
	app.Set(key, val)
	ret := app.Get(key)
	if ret != val {
		t.Error("can't get value")
	}
}

func TestApplicationEnable(t *testing.T) {
	app := NewApplication()
	key := "key"
	if x := app.Enable(key).Get(key); x != true {
		t.Error("can't enable")
	}
}

func TestApplicationEnabledd(t *testing.T) {
	app := NewApplication()
	key := "key"
	if x := app.Enable(key).Enabled(key); x != true {
		t.Error("we called enable but the key isn't enabled")
	}
	if x := app.Enabled("other"); x != false {
		t.Error("we do not have a value so it should be false")
	}
}

func TestApplicationDisable(t *testing.T) {
	app := NewApplication()
	key := "key"
	if x := app.Disable(key).Get(key); x != false {
		t.Error("disabling a key should have a value of false")
	}
}

func TestApplicationDisabled(t *testing.T) {
	app := NewApplication()
	key := "key"
	if x := app.Disable(key).Disabled(key); x != true {
		t.Error("checking if a disabled value should be false")
	}
	if x := app.Disabled("other"); x != false {
		t.Error("checking if something is disable should return true")
	}
}

func TestApplicationSetEngine(t *testing.T) {
	app := NewApplication()
	key := "html"
	engine := new(Fn)
	app.SetEngine(key, engine)
	if x := app.engines["."+key]; x != engine {
		t.Error("couldn't set the engine")
	}
}

func TestApplicationGetEngine(t *testing.T) {
	app := NewApplication()
	key := "html"
	engine := new(Fn)
	if x := app.SetEngine(key, engine).Engine(key); x != engine {
		t.Error("after setting an engine we must be able to get it")
	}
}

func TestApplicationRouter(t *testing.T) {
	app := NewApplication()
	if x := app.Router(); x == nil {
		t.Error("the router must be set")
	}
}

func TestApplicationUse(t *testing.T) {
	app := NewApplication()
	fn := new(Fn)
	app.Use(nil, fn)
	var found bool = false
	for _, x := range app.Router().stack {
		if &x.handle == fn {
			found = true
		}
	}
	if !found {
		t.Error("the fn should have been delegated to Router#use")
	}
}
