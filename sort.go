package sort

func main() {

}

func swap(intArr []int, i, j int) []int {
	temp := intArr[i]
	intArr[i] = intArr[j]
	intArr[j] = temp
	return intArr
}

func swapAdjecent(intitArr []int, i int) []int {
	for index := i; index > 0; index-- {
		if intitArr[index-1] > intitArr[index] {
			swap(intitArr, index, index-1)
		}
	}
	return intitArr
}

func sortInsertion(intitArr []int) []int {
	for i := range intitArr {
		swapAdjecent(intitArr, i)
	}
	return intitArr
}
