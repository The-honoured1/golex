package registry

import "errors"

type Processor func(string) (string, error)

var registry = map[string]Processor{}

func Register(name string, p Processor) {
	registry[name] = p
}

func Get(name string) (Processor, error) {
	p, ok := registry[name]
	if !ok {
		return nil, errors.New("processor not found")
	}
	return p, nil
}
