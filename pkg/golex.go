package pkg

import (
	"github.com/The-honoured1/golex/internal/core"
	"github.com/The-honoured1/golex/internal/pipeline"
)

type Engine struct {
	p *pipeline.Pipeline
}

func NewEngine(steps ...core.Processor) *Engine {
	return &Engine{
		p: pipeline.New(steps...),
	}
}

func (e *Engine) Process(text string) (string, error) {
	return e.p.Execute(text)
}
