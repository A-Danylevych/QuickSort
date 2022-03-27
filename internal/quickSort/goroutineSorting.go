package quickSort

const pivotIndex = 0

func GoroutineSorting(array []int, goroutinesCount int) []int {
	arrayCopy := make([]int, 0, len(array))
	copy(arrayCopy, array)
	return goroutineSorting(array, goroutinesCount)
}

func goroutineSorting(array []int, goroutinesCount int) []int {
	dividedArray := divideArray(array, goroutinesCount)
	size := len(array)
	sortedArray := make([]int, 0, size)

	communicationChanMap := make(map[int]chan []int)
	resultChanMap := make(map[int]chan []int)
	for index := 0; index < goroutinesCount; index++ {
		communicationChanMap[index] = make(chan []int, goroutinesCount-1)
		resultChanMap[index] = make(chan []int)
	}

	go transmitter(dividedArray[pivotIndex], goroutinesCount, size, pivotIndex, communicationChanMap, resultChanMap)

	for index := 1; index < goroutinesCount; index++ {
		go receiver(dividedArray[index], goroutinesCount, size, index, communicationChanMap, resultChanMap)
	}

	for index := 0; index < goroutinesCount; index++ {
		sortedArray = append(sortedArray, <-resultChanMap[index]...)
	}

	return sortedArray
}

func receiver(array []int, goroutinesCount, arraySize, id int, communicationChanMap,
	resultChanMap map[int]chan []int) {

	array = SequentialSort(array)

	localPivots := calculateLocalPivots(array, goroutinesCount, arraySize)
	sendLocalPivots(localPivots, communicationChanMap)
	pivots := getGlobalPivots(id, communicationChanMap)
	dividedArray := divideArrayByPivots(array, pivots, goroutinesCount)
	remain := sendParts(dividedArray, id, goroutinesCount, communicationChanMap)
	remain = append(remain, getParts(arraySize, id, goroutinesCount, communicationChanMap)...)
	sorted := SequentialSort(remain)
	resultChanMap[id] <- sorted
}

func transmitter(array []int, goroutinesCount, arraySize, id int, communicationChanMap,
	resultChanMap map[int]chan []int) {

	array = SequentialSort(array)

	localPivots := calculateLocalPivots(array, goroutinesCount, arraySize)
	localPivots = getLocalPivots(localPivots, goroutinesCount, communicationChanMap)

	localPivots = SequentialSort(localPivots)
	pivots := sendGlobalPivots(localPivots, communicationChanMap, goroutinesCount)
	dividedArray := divideArrayByPivots(array, pivots, goroutinesCount)
	remain := sendParts(dividedArray, id, goroutinesCount, communicationChanMap)
	remain = append(remain, getParts(arraySize, id, goroutinesCount, communicationChanMap)...)
	remain = SequentialSort(remain)
	resultChanMap[id] <- remain
}

func calculateLocalPivots(array []int, goroutinesCount, size int) []int {
	localPivots := make([]int, goroutinesCount)
	multiplier := size / (goroutinesCount * goroutinesCount)
	for index := 0; index < goroutinesCount; index++ {
		localPivots[index] = array[index*multiplier]
	}
	return localPivots
}

func getLocalPivots(pivots []int, goroutinesCount int, communicationChanMap map[int]chan []int) []int {
	for index := 1; index < goroutinesCount; index++ {
		pivots = append(pivots, <-communicationChanMap[pivotIndex]...)
	}
	return pivots
}

func sendLocalPivots(array []int, communicationChanMap map[int]chan []int) {
	communicationChanMap[pivotIndex] <- array
}

func sendGlobalPivots(pivots []int, communicationChanMap map[int]chan []int, goroutinesCount int) []int {
	globalPivots := make([]int, goroutinesCount-1)
	multiplier := 1
	addition := goroutinesCount/2 - 1
	for index := 0; index < goroutinesCount-1; index++ {
		globalPivots[index] = pivots[multiplier*goroutinesCount+addition]
		multiplier++
	}
	for index := 1; index < goroutinesCount; index++ {
		communicationChanMap[index] <- globalPivots
	}

	return globalPivots
}

func getGlobalPivots(id int, communicationChanMap map[int]chan []int) []int {
	array := <-communicationChanMap[id]
	return array
}

func divideArrayByPivots(array, pivots []int, goroutinesCount int) [][]int {
	dividedArray := make([][]int, goroutinesCount)
	for _, value := range array {
		dividedArray = aadValueToArray(dividedArray, pivots, value)
	}
	return dividedArray
}

func aadValueToArray(array [][]int, pivots []int, element int) [][]int {
	index := 0
	for _, value := range pivots {
		if element > value {
			index++
		}
	}
	array[index] = append(array[index], element)

	return array
}

func sendParts(array [][]int, id, goroutinesCount int, communicationChanMap map[int]chan []int) []int {
	for index := 0; index < goroutinesCount; index++ {
		if index == id {
			continue
		}
		communicationChanMap[index] <- array[index]
	}
	return array[id]
}

func getParts(size, id, goroutinesCount int, communicationChanMap map[int]chan []int) []int {
	array := make([]int, 0, size/goroutinesCount)
	for index := 0; index < goroutinesCount; index++ {
		if index == id {
			continue
		}
		array = append(array, <-communicationChanMap[id]...)
	}
	return array
}

func divideArray(array []int, goroutinesCount int) [][]int {
	arraySize := len(array)
	endIndexes := make([]int, goroutinesCount)
	elementsPerArray := arraySize / goroutinesCount
	extraElements := arraySize % goroutinesCount

	elementsCount := elementsPerArray

	for i := 0; i < goroutinesCount; i++ {
		if extraElements != 0 {
			elementsCount++
			extraElements--
		}
		endIndexes[i] = elementsCount
		elementsCount += elementsPerArray
	}

	dividedArray := make([][]int, 0, goroutinesCount)

	lastIndex := 0

	for index := 0; index < goroutinesCount; index++ {
		subArray := make([]int, endIndexes[index]-lastIndex)
		copy(subArray, array[lastIndex:endIndexes[index]])
		dividedArray = append(dividedArray, subArray)
		lastIndex = endIndexes[index]
	}

	return dividedArray
}
