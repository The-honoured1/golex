package pipeline

type Executor struct {
	pipeline *Pipeline
}

func NewExecutor(p *Pipeline) *Executor {
	return &Executor{pipeline: p}
}

func (e *Executor) RunBatch(inputs []string) ([]string, error) {
	results := make([]string, len(inputs))

	for i, in := range inputs {
		out, err := e.pipeline.Execute(in)
		if err != nil {
			return nil, err
		}
		results[i] = out
	}

	return results, nil
}
