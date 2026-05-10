package pipeline

import (
	"sync"

	"github.com/The-honoured1/golex/internal/core"
)

type WorkerPool struct {
	workers int
	p       *Pipeline
}

func NewWorkerPool(workers int, p *Pipeline) *WorkerPool {
	return &WorkerPool{workers: workers, p: p}
}

func (w *WorkerPool) Run(inputs []string) []core.Result {
	jobs := make(chan string)
	results := make(chan core.Result)

	var wg sync.WaitGroup

	for i := 0; i < w.workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for input := range jobs {
				out, err := w.p.Execute(input)
				results <- core.Result{Text: out, Err: err}
			}
		}()
	}

	go func() {
		for _, in := range inputs {
			jobs <- in
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var out []core.Result
	for r := range results {
		out = append(out, r)
	}

	return out
}
