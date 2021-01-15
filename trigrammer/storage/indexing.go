package storage

import (
	"sync"
)

var (
	// forward and reverse word indices
	wordIndex        = make(map[string]uint32)  // 'word' to 'id' index
	reverseWordIndex = make(map[uint32]*string) // 'id' to 'word' index
	wm               sync.Mutex
)

func wordToIndex(word string) uint32 {
	wm.Lock()
	defer wm.Unlock()
	if idx, ok := wordIndex[word]; ok {
		return idx
	}
	wordIndex[word] = uint32(len(wordIndex) + 1)
	reverseWordIndex[wordIndex[word]] = &word
	return wordIndex[word]
}

func wordByIndex(index uint32) string {
	wm.Lock()
	defer wm.Unlock()
	if word, ok := reverseWordIndex[index]; ok {
		return *word
	}
	// If there's no word by the index, it means something's wrong with our code flow
	// This should never happen
	panic("possible bug! word not found")
}
