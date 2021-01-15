package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordToIndex(t *testing.T) {
	word := "foo bar"
	assert.Equal(t, word, wordByIndex(wordToIndex(word)))
}
