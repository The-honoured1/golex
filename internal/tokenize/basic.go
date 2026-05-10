package tokenize

import "strings"

type Basic struct{}

func (t Basic) Process(input string) (string, error) {
	return strings.Join(strings.Fields(input), " "), nil
}
