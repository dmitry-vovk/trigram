package storage

import (
	"sync"
)

var (
	// forward and reverse word indices
	wordIndex        = make(map[string]uint32) // 'word' to 'id' index
	reverseWordIndex []*string                 // 'id' to 'word' index
	wm               sync.Mutex
)

func wordToIndex(word string) uint32 {
	wm.Lock()
	defer wm.Unlock()
	if idx, ok := wordIndex[word]; ok {
		return idx
	}
	wordIndex[word] = uint32(len(wordIndex) + 1)
	reverseWordIndex = append(reverseWordIndex, &word)
	return wordIndex[word]
}

func wordByIndex(index uint32) string {
	return *reverseWordIndex[index-1]
}
