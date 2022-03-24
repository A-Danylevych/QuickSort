package tester

import (
	"QuickSort/internal/generator"
	"QuickSort/internal/output"
	"QuickSort/internal/quickSort"
	"fmt"
	"time"
)

type Tester struct {
	testData          [][]int
	initialSize       int
	endSize           int
	initialGoroutines int
	endGoroutines     int
	dataCount         int
	linearResults     results
	goroutinesResults results
	sizeCount         int
}

func NewTester(initialSize, endSize, initialGoroutines, endGoroutines, dataCount int) *Tester {

	sizeCount := (endSize - initialSize) / initialSize
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
	}
}

func (t Tester) GenerateRandomData(elementRange int) {
	currentSize := t.initialSize
	for index := range t.testData {
		t.testData[index] = generator.GenerateRandomArray(currentSize, elementRange)
		if (index+1)%t.dataCount == 0 {
			currentSize += t.initialSize
		}
	}
}

func (t Tester) GeneratePermData() {
	currentSize := t.initialSize
	for index := range t.testData {
		t.testData[index] = generator.GeneratePermArray(currentSize)
		if (index+1)%t.dataCount == 0 {
			currentSize += t.initialSize
		}
	}
}

func (t Tester) SetData(data []int) {
	t.testData = append(t.testData, data)
}

func (t Tester) TestLinear() {

	t.linearResults = *newResults(t.sizeCount, t.dataCount)

	for _, value := range t.testData {
		start := time.Now()
		sorted := quickSort.LinearSort(value)
		elapsed := time.Since(start)
		t.linearResults.addResult(value, sorted, int(elapsed), 0)
	}
}

func (t Tester) DisplayStats() {
	table := output.NewTable(t.initialGoroutines, t.endGoroutines)
	for i := t.initialSize; i < t.endSize; i += t.initialSize {
		table.AddIntElement(i)
		table.AddDoubleElement(t.linearResults.getTime(i, 0))
		for j := t.initialGoroutines; j < t.endGoroutines; j++ {
			table.AddDoubleElement(t.goroutinesResults.getTime(i, j))
		}
		table.AddRow()
	}
	table.Render()
}

func (t Tester) Display() {
	for i := t.initialSize; i < t.endSize; i += t.initialSize {
		display(t.linearResults.getResult(i))
		//display(t.goroutinesResults.getResult(i))
	}
}

func display(result result) {
	for index, value := range result.time {
		fmt.Printf("Size - %d, Goroutines - %d\n", result.Size, result.GoroutinesCount[index])
		fmt.Printf("Initial array: %v\n", result.InitArray)
		fmt.Printf("Sorted array: %v\n", result.SortedArray)
		fmt.Printf("Time elepsed= %d", value)
	}
}
