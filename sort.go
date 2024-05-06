package sort

func main() {

}

func swap(intArr []int, i, j int) []int {
	temp := intArr[i]
	intArr[i] = intArr[j]
	intArr[j] = temp
	return intArr
}

func sortInsertion(intitArr []int) []int {
	if intitArr[0] > intitArr[1] {
		swap(intitArr, 1, 0)
	}
	return intitArr
}
