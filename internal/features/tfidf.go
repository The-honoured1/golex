package features

import (
	"math"
	"strings"
)

type TFIDF struct {
	docs [][]string
}

func NewTFIDF(docs []string) *TFIDF {
	processed := make([][]string, len(docs))

	for i, d := range docs {
		processed[i] = strings.Fields(d)
	}

	return &TFIDF{docs: processed}
}

func tf(term string, doc []string) float64 {
	count := 0
	for _, w := range doc {
		if w == term {
			count++
		}
	}
	return float64(count) / float64(len(doc))
}

func idf(term string, docs [][]string) float64 {
	count := 0

	for _, doc := range docs {
		for _, w := range doc {
			if w == term {
				count++
				break
			}
		}
	}

	return math.Log(float64(len(docs)) / (1 + float64(count)))
}

func (t *TFIDF) Vectorize(doc string) map[string]float64 {
	words := strings.Fields(doc)

	vec := make(map[string]float64)

	for _, w := range words {
		vec[w] = tf(w, words) * idf(w, t.docs)
	}

	return vec
}
