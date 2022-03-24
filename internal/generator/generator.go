package generator

import (
	"math/rand"
	"time"
)

func GenerateArray(size int) []int {
	rand.Seed(time.Now().Unix())

	var array = rand.Perm(size)
	return array
}
