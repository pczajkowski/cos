package cos

import (
	"strings"
	"testing"
	"unicode/utf8"
  "fmt"
)

func checkChunks(text string, size int) (*chunkOfSize, int, error) {
  chunk := NewChunkOfSize(testText, size)
	count := 0

	for {
		text := chunk.Next()
		if text == "" {
			break
		}

		if utf8.RuneCountInString(text) > size {
      return chunk, -1, fmt.Errorf("'%s'\nis longer than %d", text, size)
		}

		count++
	}

  return chunk, count, nil
}

func Test4ChunksOf100(t *testing.T) {
	size := 100
  expectedChunks := 4
  
	chunk, count, err := checkChunks(testText, size)
  if err != nil {
    t.Fatal(err)
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

func TestTextShorterThanLimit(t *testing.T) {
	size := 400
  expectedChunks := 1

  chunk, count, err := checkChunks(testText, size)
  if err != nil {
    t.Fatal(err)
  }
  
	if count != expectedChunks {
		t.Fatal("There should be", expectedChunks, "chunks, but have", count)
	}

	if !chunk.Success() {
		t.Fatal("There were errors:\n", strings.Join(chunk.GetErrors(), "\n"))
	}
}

func checkWord(word string) bool {
  for i := range validWords {
    if word == validWords[i] {
      return true
    }
  }

  return false
}

func TestWordsNotCut(t *testing.T){
  size := 12
  chunk := NewChunkOfSize(shorterText, size)

	for {
		text := chunk.Next()
		if text == "" {
			break
		}

		words := strings.Split(text, " ")
    for i := range words {
      if !checkWord(words[i]) {
        t.Fatal(words[i], "isn't a valid word!")
      }
    }
	}
}