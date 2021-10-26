package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU(4)
	v, ok := lru.Get(100)
	assert.Equal(t, -1, v)
	assert.Equal(t, false, ok)
	lru.Put(10, 0)
	lru.Put(11, 1)
	lru.Put(12, 2)
	// 0->1->2(最新)
	v, ok = lru.Get(12)
	assert.Equal(t, 2, v)
	assert.Equal(t, true, ok)
	//
	lru.Put(13, 3)
	lru.Put(14, 4)
	// 1->2->3->4
	v, ok = lru.Get(0)
	assert.Equal(t, -1, v)
	assert.Equal(t, false, ok)
	lru.Put(11, 1)
	lru.Put(11, 1)
	lru.Put(11, 1)
	lru.Put(11, 1)
}
