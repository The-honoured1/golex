package tokenize

import (
	"regexp"
	"strings"
)

type Regex struct{}

func (t Regex) Tokenize(text string) []string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	clean := re.ReplaceAllString(text, " ")
	return strings.Fields(clean)
}
