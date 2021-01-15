package storage

import (
	"testing"

	"github.com/dmitry-vovk/trigram/trigrammer/trigram"
	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	AddTrigram(&trigram.Trigram{Word1: "to", Word2: "be", Word3: "or"})
	AddTrigram(&trigram.Trigram{Word1: "to", Word2: "be", Word3: "or"}) // duplicate
	AddTrigram(&trigram.Trigram{Word1: "or", Word2: "not", Word3: "to"})
	AddTrigram(&trigram.Trigram{Word1: "be", Word2: "or", Word3: "not"})
	tr := &trigram.Trigram{Word1: "to", Word2: "be", Word3: "or"}
	sentence := trigram.NewSentence(tr)
	for {
		tr = GetMatchingTrigram(tr)
		if tr == nil {
			break
		}
		sentence.Add(tr)
	}
	assert.Equal(t, "to be or not to", sentence.String())
}
