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

func TestAdjCompareSwap(t *testing.T) {

	intitArr := []int{1, 2}
	var sortArr []int

	sortArr = adjCompareSwap(intitArr, 1)
	require.Equal(t, []int{1, 2}, sortArr)

	intitArr = []int{2, 1}
	sortArr = adjCompareSwap(intitArr, 1)
	require.Equal(t, []int{1, 2}, sortArr)

	intitArr = []int{3, 1}
	sortArr = adjCompareSwap(intitArr, 1)
	require.Equal(t, []int{1, 3}, sortArr)
}

func TestSwapAdjecent(t *testing.T) {
	var sortArr []int
	intitArr := []int{2, 1}

	sortArr = swapAdjecent(intitArr, 1)
	require.Equal(t, []int{1, 2}, sortArr)

	intitArr = []int{3, 1}
	sortArr = swapAdjecent(intitArr, 1)
	require.Equal(t, []int{1, 3}, sortArr)

	intitArr = []int{1, 3, 2}
	sortArr = swapAdjecent(intitArr, 2)
	require.Equal(t, []int{1, 2, 3}, sortArr)
}

func TestSortInsertion(t *testing.T) {
	intitArr := []int{1, 2}
	var sortArr []int

	sortArr = sortInsertion(intitArr)
	require.Equal(t, sortArr, []int{1, 2})

	intitArr = []int{2, 1}
	sortArr = sortInsertion(intitArr)
	require.Equal(t, sortArr, []int{1, 2})

	intitArr = []int{3, 2, 1}
	sortArr = sortInsertion(intitArr)
	require.Equal(t, []int{1, 2, 3}, sortArr)
}
