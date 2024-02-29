package leetcodelearn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDiff1(t *testing.T) {

	r := maxDiff(123456)
	assert.Equal(t, 820000, r, "应该是820000")

}

func TestMaxDiff2(t *testing.T) {

	r1 := maxDiff(10000)
	assert.Equal(t, 80000, r1, "应该是80000")
}

func TestMaxDiff3(t *testing.T) {

	r1 := maxDiff(555)
	assert.Equal(t, 888, r1, "应该是888")
}

func TestMaxDiff4(t *testing.T) {

	r1 := maxDiff(9)
	assert.Equal(t, 8, r1, "应该是8")
}

func TestMaxDiff5(t *testing.T) {

	r1 := maxDiff(111)
	assert.Equal(t, 888, r1, "应该是888")
}
