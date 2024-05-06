package sort

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwap(t *testing.T) {
	initArr := []int{2, 1}
	var swapArr []int

	swapArr = swap(initArr)

	require.Equal(t, swapArr, []int{1, 2})

}

func TestSortInsertion(t *testing.T) {
	intitArr := []int{1, 2}
	var sortArr []int

	sortArr = sortInsertion(intitArr)
	require.Equal(t, sortArr, []int{1, 2})
}
