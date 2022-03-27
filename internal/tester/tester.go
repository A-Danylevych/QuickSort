package tester

import (
	"QuickSort/internal/generator"
	"QuickSort/internal/output"
	"QuickSort/internal/quickSort"
	"fmt"
	"runtime"
	"time"
)

type Tester struct {
	testData          [][]int
	initialSize       int
	endSize           int
	initialGoroutines int
	endGoroutines     int
	dataCount         int
	sequentialResults results
	concurrentResult  results
	parallelResult    results
	sizeCount         int
	goStep            int
}

func NewTester(initialSize, endSize, initialGoroutines, endGoroutines, dataCount, goStep int) *Tester {

	sizeCount := (endSize-initialSize)/initialSize + 1
	iterations := dataCount * sizeCount
	testData := make([][]int, 0, iterations)

	return &Tester{
		testData:          testData,
		initialSize:       initialSize,
		endSize:           endSize,
		initialGoroutines: initialGoroutines,
		endGoroutines:     endGoroutines,
		dataCount:         dataCount,
		sizeCount:         sizeCount,
		goStep:            goStep,
	}
}

func (t *Tester) GenerateRandomData(elementRange int) {
	currentSize := t.initialSize
	for index := 0; index < cap(t.testData); index++ {
		t.testData = append(t.testData, generator.GenerateRandomArray(currentSize, elementRange))
		if (index+1)%t.dataCount == 0 {
			currentSize += t.initialSize
		}
	}
}

func (t *Tester) GeneratePermData() {
	currentSize := t.initialSize
	for index := 0; index < cap(t.testData); index++ {
		t.testData = append(t.testData, generator.GeneratePermArray(currentSize))
		if (index+1)%t.dataCount == 0 {
			currentSize += t.initialSize
		}
	}
}

func (t *Tester) SetData(data []int) {
	t.testData = append(t.testData, data)
	t.initialSize = len(data)
	t.endSize = len(data)
}

func (t *Tester) TestSequential() {

	t.sequentialResults = *newResults(t.sizeCount, t.dataCount)

	for _, value := range t.testData {
		start := time.Now()
		sorted := quickSort.SequentialSort(value)
		elapsed := time.Since(start).Microseconds()
		t.sequentialResults.addResult(value, sorted, elapsed, 0)
	}
}

func (t *Tester) TestConcurrent() {
	t.concurrentResult = *newResults(t.sizeCount, t.dataCount)

	for _, value := range t.testData {
		for count := t.initialGoroutines; count <= t.endGoroutines; count += t.goStep {
			start := time.Now()
			sorted := quickSort.GoroutineSorting(value, count)
			elapsed := time.Since(start).Microseconds()
			t.concurrentResult.addResult(value, sorted, elapsed, count)
		}

	}
}

func (t *Tester) TestParallel() {
	t.parallelResult = *newResults(t.sizeCount, t.dataCount)

	for _, value := range t.testData {
		for count := t.initialGoroutines; count <= t.endGoroutines; count += t.goStep {
			runtime.GOMAXPROCS(count)
			start := time.Now()
			sorted := quickSort.GoroutineSorting(value, count)
			elapsed := time.Since(start).Microseconds()
			t.parallelResult.addResult(value, sorted, elapsed, count)
		}

	}
}

func (t *Tester) DisplayStats() {
	table := output.NewTable(t.initialGoroutines, t.endGoroutines, t.goStep)
	for i := t.initialSize; i <= t.endSize; i += t.initialSize {
		table.AddIntElement(i)
		sequentialTime := t.sequentialResults.getTime(i, 0)
		table.AddDoubleElement(sequentialTime)
		for j := t.initialGoroutines; j <= t.endGoroutines; j += t.goStep {
			goroutineTime := t.concurrentResult.getTime(i, j)
			table.AddDoubleElement(goroutineTime)
			table.AddDoubleElement(sequentialTime / goroutineTime)
			goroutineTime = t.parallelResult.getTime(i, j)
			table.AddDoubleElement(goroutineTime)
			table.AddDoubleElement(sequentialTime / goroutineTime)
		}
		table.AddRow()
	}
	table.Render()
}

func (t *Tester) Display() {
	for i := t.initialSize; i <= t.endSize; i += t.initialSize {
		fmt.Printf("Sequential\n")
		display(t.sequentialResults.getResult(i))
		fmt.Printf("Concurrent\n")
		display(t.concurrentResult.getResult(i))
		fmt.Printf("Paralel\n")
		display(t.parallelResult.getResult(i))
	}
}

func display(result result) {
	for index, value := range result.time {
		fmt.Printf("Size - %d, Goroutines - %d\n", result.Size, result.GoroutinesCount[index])
		fmt.Printf("Initial array: %v\n", result.InitArray[index])
		fmt.Printf("Sorted array: %v\n", result.SortedArray[index])
		fmt.Printf("Time elepsed= %d\n", value)
	}
}
