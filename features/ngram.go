package features

import "strings"

type NGram struct {
	N int
}

func (n NGram) Process(input string) (string, error) {
	words := strings.Fields(input)

	if len(words) < n.N {
		return input, nil
	}

	ngrams := []string{}
	for i := 0; i <= len(words)-n.N; i++ {
		ngrams = append(ngrams, strings.Join(words[i:i+n.N], "_"))
	}

	return strings.Join(ngrams, " "), nil
}
