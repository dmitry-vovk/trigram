package storage

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// thirdWords contains the list of third words in a trigram
type thirdWords struct {
	m       sync.RWMutex
	number  int
	indices map[uint32]int
}

// idx corresponds to the two words tying two trigrams together
type idx struct {
	w1 uint32
	w2 uint32
}

func init() { rand.Seed(time.Now().UnixNano()) }

// add inserts another index into thirdWords structure
func (w *thirdWords) add(index uint32) *thirdWords {
	w.m.Lock()
	defer w.m.Unlock()
	if _, ok := w.indices[index]; ok {
		w.indices[index]++
	} else {
		w.indices[index] = 1
	}
	w.number++
	return w
}

// random returns a third word randomly and proportionally to usage frequency
func (w *thirdWords) random() uint32 {
	if w.number == 0 {
		// Having a trigram without the third word should not happen
		log.Fatalf("Fatal error! Calling random on empty index.")
	}
	w.m.RLock()
	defer w.m.RUnlock()
	if len(w.indices) == 1 {
		for index := range w.indices {
			return index
		}
	}
	n := rand.Intn(w.number-1) + 1
	m := 0
	var index uint32
	var offset int
	for index, offset = range w.indices {
		if n > m && n < m+offset {
			return index
		}
		m += offset
	}
	return index
}
