package gopress

type Layer struct {
	keys   []string
	params map[string]interface{}
	regexp interface{}
	handle interface{}
}

func NewLayer(path string, options map[string]interface{}, handle *interface{}) *Layer {
	layer := new(Layer)
	layer.keys = make([]string, 1)
	return layer
}
