package main

import (
	"fmt"

	"github.com/The-honoured1/lexigo/core"
	"github.com/The-honoured1/lexigo/pipeline"
	"github.com/The-honoured1/lexigo/tokenize"
)

func main() {
	p := pipeline.NewPipeline(
		core.Normalize{},
		tokenize.BasicTokenizer{},
	)

	out, err := p.Run("Hello WORLD from Lexigo NLP Engine")
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
