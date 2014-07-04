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

func TestEnable(t *testing.T) {
	app := Gopress()
	key := "key"
	if x := app.Enable(key).Get(key); x != true {
		t.Error("can't enable")
	}
}

func TestEnabledd(t *testing.T) {
	app := Gopress()
	key := "key"
	if x := app.Enable(key).Enabled(key); x != true {
		t.Error("we called enable but the key isn't enabled")
	}
	if x := app.Enabled("other"); x != false {
		t.Error("we do not have a value so it should be false")
	}
}

func TestDisable(t *testing.T) {
	app := Gopress()
	key := "key"
	if x := app.Disable(key).Get(key); x != false {
		t.Error("disabling a key should have a value of false")
	}
}

func TestDisabled(t *testing.T) {
	app := Gopress()
	key := "key"
	if x := app.Disable(key).Disabled(key); x != true {
		t.Error("checking if a disabled value should be false")
	}
	if x := app.Disabled("other"); x != false {
		t.Error("checking if something is disable should return true")
	}
}
