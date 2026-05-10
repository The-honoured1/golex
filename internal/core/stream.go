package core

import (
	"bufio"
	"strings"
)

type StreamProcessor struct {
	ChunkSize int
}

func (s StreamProcessor) Process(input string) ([]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	chunks := []string{}
	buffer := []string{}

	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())

		if len(buffer) >= s.ChunkSize {
			chunks = append(chunks, strings.Join(buffer, " "))
			buffer = buffer[:0]
		}
	}

	if len(buffer) > 0 {
		chunks = append(chunks, strings.Join(buffer, " "))
	}

	return chunks, nil
}
