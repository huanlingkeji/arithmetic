package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU(4)
	v, ok := lru.Get(100)
	assert.Equal(t, nil, v)
	assert.Equal(t, false, ok)
	lru.Put(10, 10)
	lru.Put(11, 11)
	lru.Put(12, 12)
	v, ok = lru.Get(12)
	assert.Equal(t, 12, v)
	assert.Equal(t, true, ok)
}
