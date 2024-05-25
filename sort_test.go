package sort

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwap(t *testing.T) {
	cases := []struct {
		b []int
		a []int
		i int
		j int
	}{
		{b: []int{2, 1}, a: []int{1, 2}, i: 1, j: 0},
		{b: []int{3, 2}, a: []int{2, 3}, i: 1, j: 0},
		{b: []int{3, 2, 1}, a: []int{1, 2, 3}, i: 2, j: 0},
	}

	for r, c := range cases {
		swap(c.b, c.i, c.j)
		require.Equal(t, c.a, c.b)
	}
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
