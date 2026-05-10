package core

import "strings"

type Normalizer struct{}

func (n Normalizer) Process(input string) (string, error) {
	return strings.ToLower(strings.TrimSpace(input)), nil
}
