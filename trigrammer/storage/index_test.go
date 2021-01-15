package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	tw := thirdWords{
		indices: make(map[uint32]int),
	}
	assert.Panics(t, func() {
		tw.random()
	})
	tw.add(1)
	tw.add(2)
	tw.add(2)
	assert.Equal(t, 3, tw.number)
	assert.Equal(t, 2, len(tw.indices))
	for i := 0; i < 100; i++ {
		r := tw.random()
		assert.True(t, r == 1 || r == 2, "Unexpected random %d", r)
	}
}
