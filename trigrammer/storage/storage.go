package storage

import (
	"sync"

	"github.com/dmitry-vovk/trigram/trigrammer/trigram"
)

var (
	matches = make(map[idx]*thirdWords) // two words to the third word correspondence
	um      sync.Mutex
)

// AddTrigram adds another trigram.Trigram into 'storage'
func AddTrigram(t *trigram.Trigram) {
	um.Lock()
	index := idx{
		wordToIndex(t.Word1),
		wordToIndex(t.Word2),
	}
	if words, ok := matches[index]; ok {
		words.add(wordToIndex(t.Word3))
	} else {
		matches[index] = (&thirdWords{
			indices: make(map[uint32]int),
		}).add(wordToIndex(t.Word3))
	}
	um.Unlock()
}

// GetMatchingTrigram finds matching trigram for a given, returns nil if none match
func GetMatchingTrigram(t *trigram.Trigram) *trigram.Trigram {
	um.Lock()
	defer um.Unlock()
	index := idx{wordToIndex(t.Word2), wordToIndex(t.Word3)}
	if words, ok := matches[index]; ok {
		return &trigram.Trigram{
			Word1: t.Word2,
			Word2: t.Word3,
			Word3: wordByIndex(words.random()),
		}
	}
	return nil
}
