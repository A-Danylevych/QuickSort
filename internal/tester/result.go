package tester

type result struct {
	InitArray       [][]int
	SortedArray     [][]int
	time            []int
	GoroutinesCount []int
	Size            int
}

func newResult(initArray, sortedArray []int, time, goroutinesCount, size, resultsCount int) *result {
	init := make([][]int, 0, resultsCount)
	init = append(init, initArray)
	sorted := make([][]int, 0, resultsCount)
	sorted = append(sorted, sortedArray)
	timeArray := make([]int, 0, resultsCount)
	timeArray = append(timeArray, time)
	goroutines := make([]int, 0, resultsCount)
	goroutines = append(goroutines, goroutinesCount)
	return &result{
		InitArray:       init,
		SortedArray:     sorted,
		time:            timeArray,
		GoroutinesCount: goroutines,
		Size:            size,
	}
}

func (r *result) addResult(initArray, sortedArray []int, time, goroutinesCount int) {
	r.InitArray = append(r.InitArray, initArray)
	r.SortedArray = append(r.SortedArray, sortedArray)
	r.time = append(r.time, time)
	r.GoroutinesCount = append(r.GoroutinesCount, goroutinesCount)
}

func (r *result) getTime(goroutinesCount int) float64 {
	var time float64
	count := 0
	for index, value := range r.time {
		if r.GoroutinesCount[index] == goroutinesCount {
			time += float64(value)
			count++
		}
	}
	time /= float64(count)
	return time
}
