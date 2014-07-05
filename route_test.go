package gopress

import "testing"

func TestNewRoute(t *testing.T) {
	path := "/some/path"
	a := NewRoute(path)
	if a == nil {
		t.Error("NewRoute must make a new route")
	}
	if a.path != path {
		t.Error("path must be what is configured")
	}
	b := NewRoute("")
	if b.path != "/" {
		t.Error("default path must be set")
	}
}
