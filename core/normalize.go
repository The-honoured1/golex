package core

import "strings"

type Normalize struct{}

func (n Normalize) Process(input string) (string, error) {
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)
	return input, nil
}
