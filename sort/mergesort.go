package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	Lindex, Rindex := 0, 0
	l, r := len(left), len(right)
	for Lindex < l && Rindex < r {
		if left[Lindex] > right[Rindex] {
			result = append(result, right[Rindex])
			Rindex++
			continue
		}
		result = append(result, left[Lindex])
		Lindex++
	}
	result = append(result, right[Rindex:]...)
	result = append(result, left[Lindex:]...)
	return result
}

func main() {

	b := []int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		a := rand.Intn(88)+10+i
		b = append(b, a)
	}
	fmt.Println(mergeSort(b))

}
