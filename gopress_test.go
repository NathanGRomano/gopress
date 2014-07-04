package gopress

import "testing"

func TestGopress(t *testing.T) {
	var app = Gopress()
	if app == nil {
		t.Error("App should be defined")
	}
}
