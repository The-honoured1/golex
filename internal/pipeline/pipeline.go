package pipeline

import "github.com/The-honoured1/golex/internal/core"

type Pipeline struct {
	steps []core.Processor
}

func New(steps ...core.Processor) *Pipeline {
	return &Pipeline{steps: steps}
}

func (p *Pipeline) Execute(input string) (string, error) {
	var err error
	out := input

	for _, step := range p.steps {
		out, err = step.Process(out)
		if err != nil {
			return "", err
		}
	}

	return out, nil
}
