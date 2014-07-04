package gopress

import "testing"

func TestGopress(t *testing.T) {
	var app = Gopress()
	if app == nil {
		t.Error("App should be defined")
	}
}

func TestUse(t *testing.T) {
	var app = Gopress()
	var fn = func() {}
	if x := app.Use(fn); x != app {
		t.Error("Expected app to return itself")
	}
	if app.fns.Len() != 1 {
		t.Error("Expected fns length to be \"1\"")
	}
}
