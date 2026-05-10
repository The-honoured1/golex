package main

import (
	"fmt"
	"os"

	"github.com/The-honoured1/golex/internal/core"
	"github.com/The-honoured1/golex/internal/pipeline"
	"github.com/The-honoured1/golex/internal/tokenize"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: lexigo \"text\"")
		return
	}

	input := os.Args[1]

	p := pipeline.New(
		core.Normalizer{},
		tokenize.Regex{},
	)

	out, err := p.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
