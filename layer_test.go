package gopress

import "testing"

func TestLayerNew(t *testing.T) {
	path := "/"
	options := make(map[string]interface{}, 3)
	options["caseSensitive"] = false
	options["strict"] = false
	options["ends"] = false
	fn := new(Fn)
	layer := NewLayer(path, options, fn, nil)
	if layer == nil {
		t.Error("must create a new layer")
	}
	if layer.path != path {
		t.Error("path must be what was assigned")
	}
	if layer.options != options {
		t.Error("options must be what was assigned")
	}
	if layer.fn != fn {
		t.Error("fn must be what was assigned")
	}
}
