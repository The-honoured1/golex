package core

import (
	"bufio"
	"strings"
)

type Stream struct {
	ChunkSize int
}

func (s Stream) Chunk(input string) []string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	var chunks []string
	var buffer []string

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

	return chunks
}
