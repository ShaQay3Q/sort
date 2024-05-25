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

	for _, c := range cases {
		swap(c.b, c.i, c.j)
		require.Equal(t, c.a, c.b)
	}
}

func TestAdjCompareSwap(t *testing.T) {

	for _, c := range []struct {
		before []int
		after  []int
		i      int
	}{
		{before: []int{1, 2}, after: []int{1, 2}, i: 1},
		{before: []int{2, 1}, after: []int{1, 2}, i: 1},
		{before: []int{3, 1}, after: []int{1, 3}, i: 1},
	} {
		adjCompareSwap(c.before, c.i)
		require.Equal(t, c.after, c.before)
	}
}

func TestSwapAdjecent(t *testing.T) {

	for _, c := range []struct {
		b []int
		a []int
		i int
	}{
		{b: []int{2, 1}, a: []int{1, 2}, i: 1},
		{b: []int{3, 1}, a: []int{1, 3}, i: 1},
		{b: []int{1, 3, 2}, a: []int{1, 2, 3}, i: 2},
	} {
		swapAdjecent(c.b, c.i)
		require.Equal(t, c.b, c.a)
	}
}

func TestSortInsertion(t *testing.T) {

	for _, c := range []struct {
		b []int
		a []int
	}{
		{b: []int{1, 2}, a: []int{1, 2}},
		{b: []int{2, 1}, a: []int{1, 2}},
		{b: []int{3, 1}, a: []int{1, 3}},
		{b: []int{3, 2, 1}, a: []int{1, 2, 3}},
	} {
		sortInsertion(c.b)
		require.Equal(t, c.b, c.a)
	}
}
