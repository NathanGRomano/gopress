package gopress

import "testing"

func TestGopress(t *testing.T) {
	var app = Gopress()
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
	app := Gopress()
	key := "key"
	val := "value"
	app.Set(key, val)
	if app.settings[key] != val {
		t.Error("can't set value")
	}
}

func TestGet(t *testing.T) {
	app := Gopress()
	key := "key"
	val := "value"
	app.Set(key, val)
	ret := app.Get(key)
	if ret != val {
		t.Error("can't get value")
	}
}
