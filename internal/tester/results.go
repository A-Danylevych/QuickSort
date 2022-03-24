package tester

type results struct {
	results      []result
	resultsCount int
}

func newResults(sizeCount, dataCount int) *results {
	resultsArray := make([]result, 0, sizeCount)
	return &results{
		results:      resultsArray,
		resultsCount: dataCount,
	}
}

func (r *results) addResult(initArray, sortedArray []int, time, goroutinesCount int) {
	size := len(initArray)
	for _, value := range r.results {
		if value.Size == size {
			value.addResult(initArray, sortedArray, time, goroutinesCount)
			return
		}
	}
	result := *newResult(initArray, sortedArray, time, goroutinesCount, size, r.resultsCount)
	r.results = append(r.results, result)
}

func (r *results) getTime(size, goroutinesCount int) float64 {
	for _, value := range r.results {
		if value.Size == size {
			return value.getTime(goroutinesCount)
		}
	}
	return 0
}

func (r *results) getResult(size int) result {
	var sizeResult result
	for _, value := range r.results {
		if value.Size == size {
			sizeResult = value
		}
	}
	return sizeResult
}
