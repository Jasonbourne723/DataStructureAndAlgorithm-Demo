package leetcodelearn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestOddNumber(t *testing.T) {
	num := largestOddNumber("239537672423884969653287101")
	assert.Equal(t, "239537672423884969653287101", num, "not equel")
}
