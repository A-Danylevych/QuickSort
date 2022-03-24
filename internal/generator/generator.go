package generator

import (
	"math/rand"
	"time"
)

func GeneratePermArray(size int) []int {
	rand.Seed(time.Now().Unix())

	var array = rand.Perm(size)
	return array
}

func GenerateRandomArray(size, elementRange int) []int {
	rand.Seed(time.Now().Unix())

	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(elementRange) - rand.Intn(elementRange)
	}
	return array
}
