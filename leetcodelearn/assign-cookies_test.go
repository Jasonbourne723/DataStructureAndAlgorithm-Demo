package leetcodelearn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignCookies(t *testing.T) {

	g := []int{1, 2, 3}
	s := []int{1, 1}

	num := findContentChildren(g, s)

	assert.Equal(t, 1, num, "num should is 1")

}

func TestAssignCookies2(t *testing.T) {

	g := []int{1, 2, 3}
	s := []int{3}

	num := findContentChildren(g, s)

	assert.Equal(t, 1, num, "num should is 1")

}
func TestAssignCookies1(t *testing.T) {

	g := []int{1, 2}
	s := []int{1, 2, 3}

	num := findContentChildren(g, s)

	assert.Equal(t, 2, num, "num should is 2")

}

func TestAssignCookies3(t *testing.T) {

	g := []int{2, 3, 4, 5, 6, 7, 8, 9}
	s := []int{3, 4, 6, 8, 1, 2, 2, 2}

	num := findContentChildren(g, s)

	assert.Equal(t, 5, num, "num should is 5")

}
