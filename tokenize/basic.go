package tokenize

import "strings"

type BasicTokenizer struct{}

func (t BasicTokenizer) Process(input string) (string, error) {
	return strings.Join(strings.Fields(input), " "), nil
}
