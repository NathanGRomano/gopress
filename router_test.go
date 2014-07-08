package gopress

import "testing"

func TestRouterNewRouter(t *testing.T) {
	a := NewRouter(nil)
	if a == nil {
		t.Error("router must not be nil")
	}
	if a.caseSensitive != false {
		t.Error("caseSensitive must be false by default")
	}
	if a.strict != false {
		t.Error("strict must be false by default")
	}
	config := map[string]bool{
		"caseSensitive": true,
		"strict":        true,
		"mergeParams":   true,
	}
	b := NewRouter(config)
	if b.caseSensitive != true {
		t.Error("caseSensitive must be true when set in the config")
	}
	if b.strict != true {
		t.Error("strict must be true when set in the config")
	}
	if b.mergeParams != true {
		t.Error("mergeParams must be true when set in the config")
	}
}

func TestRouterUse(t *testing.T) {
	a := NewRouter(nil)
	b := new(Fn)
	a.Use(nil, b)
	var found bool = false
	for _, x := range a.stack {
		if x.fn == b {
			found = true
		}
	}
	if !found {
		t.Error("when adding an Fn to a Path it must be added to the list")
	}
}
