package locde

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinLen(t *testing.T) {
	s1 := "aa"
	s2 := "bbb"
	assert.Equal(t, 3, getMinLen(s1, s2))
}
