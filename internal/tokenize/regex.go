package tokenize

import (
	"regexp"
	"strings"
)

type Regex struct{}

func (t Regex) Process(input string) (string, error) {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	clean := re.ReplaceAllString(input, "")
	return strings.Join(strings.Fields(clean), " "), nil
}
