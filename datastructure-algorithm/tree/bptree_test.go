package DataStructureAndAlgorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBPTree(t *testing.T) {

	bpTree := NewBPTree(4)

	bpTree.Set(10, "hello10")
	bpTree.Set(11, "hello11")
	bpTree.Set(12, "hello12")
	bpTree.Set(100, "hello100")
	bpTree.Set(9, "hello9")
	bpTree.Set(7, "hello7")
	bpTree.Set(19, "hello19")
	bpTree.Set(210, "hello210")
	bpTree.Set(140, "hello140")
	bpTree.Set(1, "hello1")

	val := bpTree.Get(11)

	bpTree.List()

	assert.Equal(t, "hello11", val, "is fail")

}
