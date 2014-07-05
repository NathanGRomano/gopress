package gopress

import "testing"

func TestNewRouter(t *testing.T) {
	router := Router()
	if router == nil {
		t.Error("router must not be nil")
	}
	if router.options["caseSensitive"] != false {
		t.Error("caseSensitive must be false by default")
	}
	if router.options["strict"] != false {
		t.Error("strict must be false by default")
	}
	config := map[string]bool{
		"caseSensitive": true,
		"strict":        true,
	}
	router := Router(config)
	if router.options["caseSensitive"] != true {
		t.Error("caseSensitive must be true when set in the config")
	}
	if router.options["strict"] != true {
		t.Error("strict must be true when set in the config")
	}
}
