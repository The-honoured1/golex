package tokenize

type Tokenizer interface {
	Tokenize(text string) []string
}
