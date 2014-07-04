package gopress

import "testing"

func TestGoPress(t *testing.T) {
	if app := GoPress(); app == nil {
		t.Error("GoPress() must make an Application")
	}
}

func TestGopress(t *testing.T) {
	if app := GoPress(); app == nil {
		t.Error("Gopress() must make an Application")
	}
}
