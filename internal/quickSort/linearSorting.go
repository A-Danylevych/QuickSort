package quickSort

func LinearSort(array []int) []int {
	var arrayCopy = make([]int, len(array))
	copy(arrayCopy, array)
	return linearSort(arrayCopy, 0, len(arrayCopy)-1)
}

func linearSort(array []int, lowIndex, highIndex int) []int {
	if lowIndex < highIndex {
		var pivot int
		array, pivot = partition(array, lowIndex, highIndex)
		array = linearSort(array, lowIndex, pivot-1)
		array = linearSort(array, pivot+1, highIndex)
	}
	return array
}

func partition(array []int, lowIndex, highIndex int) ([]int, int) {
	pivot := array[highIndex]
	PivotIndex := lowIndex
	for j := lowIndex; j < highIndex; j++ {
		if array[j] < pivot {
			array[PivotIndex], array[j] = array[j], array[PivotIndex]
			PivotIndex++
		}
	}
	array[PivotIndex], array[highIndex] = array[highIndex], array[PivotIndex]
	return array, PivotIndex
}
