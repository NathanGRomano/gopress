package gopress

import "testing"

func TestNewApplication(t *testing.T) {
	var app = NewApplication()
	if app == nil {
		t.Error("App should be defined")
	}
}

func TestInit(t *testing.T) {
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

func TestDefaultConfiguration(t *testing.T) {
	var app = new(Application)
	app.Init()
	app.DefaultConfiguration()
}

func TestSet(t *testing.T) {
	app := NewApplication()
	key := "key"
	val := "value"
	app.Set(key, val)
	if app.settings[key] != val {
		t.Error("can't set value")
	}
}

func TestGet(t *testing.T) {
	app := NewApplication()
	key := "key"
	val := "value"
	app.Set(key, val)
	ret := app.Get(key)
	if ret != val {
		t.Error("can't get value")
	}
}

func TestEnable(t *testing.T) {
	app := NewApplication()
	key := "key"
	if x := app.Enable(key).Get(key); x != true {
		t.Error("can't enable")
	}
}

func TestEnabledd(t *testing.T) {
	app := NewApplication()
	key := "key"
	if x := app.Enable(key).Enabled(key); x != true {
		t.Error("we called enable but the key isn't enabled")
	}
	if x := app.Enabled("other"); x != false {
		t.Error("we do not have a value so it should be false")
	}
}

func TestDisable(t *testing.T) {
	app := NewApplication()
	key := "key"
	if x := app.Disable(key).Get(key); x != false {
		t.Error("disabling a key should have a value of false")
	}
}

func TestDisabled(t *testing.T) {
	app := NewApplication()
	key := "key"
	if x := app.Disable(key).Disabled(key); x != true {
		t.Error("checking if a disabled value should be false")
	}
	if x := app.Disabled("other"); x != false {
		t.Error("checking if something is disable should return true")
	}
}

func TestSetEngine(t *testing.T) {
	app := NewApplication()
	key := "html"
	engine := new(Fn)
	app.SetEngine(key, engine)
	if x := app.engines["."+key]; x != engine {
		t.Error("couldn't set the engine")
	}
}

func TestGetEngine(t *testing.T) {
	app := NewApplication()
	key := "html"
	engine := new(Fn)
	if x := app.SetEngine(key, engine).Engine(key); x != engine {
		t.Error("after setting an engine we must be able to get it")
	}
}

func TestRouter(t *testing.T) {
	app := NewApplication()
	if x := app.Router(); x == nil {
		t.Error("the router must be set")
	}
}
