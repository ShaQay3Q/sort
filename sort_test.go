package sort

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwap(t *testing.T) {
	var swapArr []int

	swapArr = swap([]int{2, 1}, 1, 0)
	require.Equal(t, swapArr, []int{1, 2})

	swapArr = swap([]int{3, 2}, 1, 0)
	require.Equal(t, swapArr, []int{2, 3})

	swapArr = swap([]int{3, 2, 1}, 0, 2)
	require.Equal(t, swapArr, []int{1, 2, 3})
}

func TestSortInsertion(t *testing.T) {
	intitArr := []int{1, 2}
	var sortArr []int

	sortArr = sortInsertion(intitArr)
	require.Equal(t, sortArr, []int{1, 2})

	intitArr = []int{2, 1}
	sortArr = sortInsertion(intitArr)
	require.Equal(t, sortArr, []int{1, 2})
}
