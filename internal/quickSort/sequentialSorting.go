package quickSort

func SequentialSort(array []int) []int {
	var arrayCopy = make([]int, len(array))
	copy(arrayCopy, array)
	return sequentialSort(arrayCopy, 0, len(arrayCopy)-1)
}

func sequentialSort(array []int, lowIndex, highIndex int) []int {
	if lowIndex < highIndex {
		var pivot int
		array, pivot = partition(array, lowIndex, highIndex)
		array = sequentialSort(array, lowIndex, pivot-1)
		array = sequentialSort(array, pivot+1, highIndex)
	}
	return array
}

func partition(array []int, lowIndex, highIndex int) ([]int, int) {
	pivot := array[highIndex]
	pivotIndex := lowIndex
	for index := lowIndex; index < highIndex; index++ {
		if array[index] < pivot {
			array[pivotIndex], array[index] = array[index], array[pivotIndex]
			pivotIndex++
		}
	}
	array[pivotIndex], array[highIndex] = array[highIndex], array[pivotIndex]
	return array, pivotIndex
}
