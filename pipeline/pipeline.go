package pipeline

type Processor interface {
	Process(string) (string, error)
}

type Pipeline struct {
	steps []Processor
}

func NewPipeline(steps ...Processor) *Pipeline {
	return &Pipeline{steps: steps}
}

func (p *Pipeline) Run(input string) (string, error) {
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
