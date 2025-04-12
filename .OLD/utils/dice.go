package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Roll(die int) int {
	return rand.Intn(die) + 1
}

func RollMany(die, count int) []int {
	results := make([]int, count)
	for i := range results {
		results[i] = Roll(die)
	}
	return results
}
