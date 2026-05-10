package tokenize

import "regexp"

type RegexTokenizer struct{}

func (t RegexTokenizer) Process(input string) (string, error) {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return re.ReplaceAllString(input, ""), nil
}
