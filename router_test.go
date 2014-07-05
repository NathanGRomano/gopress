package gopress

import "testing"

func TestNewRouter(t *testing.T) {
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
