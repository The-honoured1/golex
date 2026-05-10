package core

type Document string

type Result struct {
	Text string
	Err  error
}

// Processor is the base contract for ALL NLP operations
type Processor interface {
	Process(input string) (string, error)
}
