package sort

func main() {

}

// swap; swaps two array indexes
func swap(intArr []int, i, j int) {
	intArr[i], intArr[j] = intArr[j], intArr[i]
}

// ajdCompareSwap: compare and swaps two adjecent indexes on an array
func adjCompareSwap(intitArr []int, i int) {
	if intitArr[i-1] > intitArr[i] {
		swap(intitArr, i, i-1)
	}
}

// swapAdjecent: do an iteration on indexes smaller than i and swap
// indexes elements when it's needed
func swapAdjecent(intitArr []int, i int) {
	for index := i; index > 0; index-- {
		adjCompareSwap(intitArr, index)
	}
}

// sortInsertion: sort an array of int using imsertion sorting method
func sortInsertion(intitArr []int) []int {
	for i := range intitArr {
		swapAdjecent(intitArr, i)
	}
	return intitArr
}
