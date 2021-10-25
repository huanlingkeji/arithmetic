package substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindow(t *testing.T) {
	for _, v := range []struct {
		S string
		T string
		O string
	}{{
		S: "ADOBECODEBANC",
		T: "ABC",
		O: "BANC",
	}} {
		assert.Equal(t, v.O, minWindow(v.S, v.T))
	}
}
