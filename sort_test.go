package sort

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwap(t *testing.T) {
	var swapArr = []int{2, 1}

	swap(swapArr, 1, 0)
	require.Equal(t, swapArr, []int{1, 2})

	swapArr = []int{3, 2}
	swap(swapArr, 1, 0)
	require.Equal(t, swapArr, []int{2, 3})

	swapArr = []int{3, 2, 1}
	swap(swapArr, 0, 2)
	require.Equal(t, swapArr, []int{1, 2, 3})
}

func TestAdjCompareSwap(t *testing.T) {

	intitArr := []int{1, 2}
	adjCompareSwap(intitArr, 1)

	require.Equal(t, []int{1, 2}, intitArr)

	intitArr = []int{2, 1}
	adjCompareSwap(intitArr, 1)

	require.Equal(t, []int{1, 2}, intitArr)

	intitArr = []int{3, 1}
	adjCompareSwap(intitArr, 1)
	require.Equal(t, []int{1, 3}, intitArr)
}

func TestSwapAdjecent(t *testing.T) {
	intitArr := []int{2, 1}

	swapAdjecent(intitArr, 1)
	require.Equal(t, []int{1, 2}, intitArr)

	intitArr = []int{3, 1}
	swapAdjecent(intitArr, 1)
	require.Equal(t, []int{1, 3}, intitArr)

	intitArr = []int{1, 3, 2}
	swapAdjecent(intitArr, 2)
	require.Equal(t, []int{1, 2, 3}, intitArr)
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
