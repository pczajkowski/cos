package cos

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func Test4ChunksOf100(t *testing.T) {
	size := 100
  expectedChunks := 4
  
	chunk := NewChunkOfSize(testText, size)
	count := 0

	for {
		text := chunk.Next()
		if text == "" {
			break
		}

		if utf8.RuneCountInString(text) > size {
			t.Fatal(text, "\nis longer than", size)
		}

		count++
	}
  
	if count != expectedChunks {
		t.Fatal("There should be", expectedChunks, "chunks, but have", count)
	}

	if !chunk.Success() {
		t.Fatal("There were errors:\n", strings.Join(chunk.GetErrors(), "\n"))
	}
}

func TestWordBiggerThanLimit(t *testing.T) {
	size := 4
  
	chunk := NewChunkOfSize(testText, size)

	text := chunk.Next()
  if text != "" {
    t.Fatal("Chunk should be empty, but is:", text)
  }

	if chunk.Success() || len(chunk.GetErrors()) == 0 {
		t.Fatal("There should be errors!")
	}
}