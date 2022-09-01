package cos

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ChunkOfSize
type ChunkOfSize struct {
	current int
	limit   int
	chunks  []string
	errors  []string
}

// NewChunkOfSize returns new instance of ChunkOfSize initialized with given text and size
func NewChunkOfSize(text string, size int) *ChunkOfSize {
	return &ChunkOfSize{
		current: 0,
		limit:   size,
		chunks:  strings.Split(text, " "),
	}
}

// Next returns next chunk of text or empty string if nothing left to process
func (c *ChunkOfSize) Next() string {
	var b strings.Builder
	for i := range c.chunks {
		l := utf8.RuneCountInString(c.chunks[i])
		if l >= c.limit {
			c.errors = append(c.errors, fmt.Sprintf("Chunk {%s} is bigger than limit %d!", c.chunks[i], c.limit))
			c.chunks = []string{}
			return ""
		}

		if l+c.current >= c.limit {
			c.current = 0
			c.chunks = c.chunks[i:]
			return b.String()
		}

		var toWrite string
		if c.current == 0 || c.chunks[i] == "\n" {
			toWrite = c.chunks[i]
		} else {
			toWrite = fmt.Sprintf(" %s", c.chunks[i])
			l++
		}

		w, err := b.WriteString(toWrite)
		if w == 0 || err != nil {
			c.errors = append(c.errors, fmt.Sprintf("Error writing: %s; %s", toWrite, err))
			continue
		}

		c.current += l
	}

	c.chunks = []string{}
	return b.String()
}

// Success returns true if there are no errors
func (c *ChunkOfSize) Success() bool {
	return len(c.errors) == 0
}

// GetErrors returns erorrs
func (c *ChunkOfSize) GetErrors() []string {
	return c.errors
}
