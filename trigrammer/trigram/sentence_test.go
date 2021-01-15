package trigram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentence(t *testing.T) {
	s := NewSentence(&Trigram{Word1: "To", Word2: "be", Word3: "or"}).
		Add(&Trigram{Word1: "be", Word2: "or", Word3: "not"}).
		Add(&Trigram{Word1: "or", Word2: "not", Word3: "to"}).
		Add(&Trigram{Word1: "not", Word2: "to", Word3: "be?"})
	assert.Equal(t, "To be or not to be?", s.String())
}
