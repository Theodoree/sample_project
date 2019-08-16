package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)



func primePalindrome(N int) int {

}

func main() {

	/*
	   f := [][]int{{2},
	       {3, 4},
	       {6, 5, 7},
	       {4, 1, 8, 3}}
	   fmt.Println(minimumTotal(f))

	   f = [][]int{{-1}, {2, 3}, {1, -1, -3}}
	   fmt.Println(minimumTotal(f))
	*/
}

// 二分查找，O(log(n / 2 + 1))
func arrangeCoins3(n int) int {
	l, r := 0, n/2+1
	for r-l > 1 {
		mid := (l + r) >> 1
		if mid*(mid+1) == 2*n {
			return mid

		} else if mid*(mid+1) <= 2*n {
			l = mid
		} else {
			r = mid
		}
	}

	if r*(r+1) == 2*n {
		return r
	}

	return l
}

func GetSlice(i int) []int {
	rand.Seed(time.Now().Unix())

	var result []int
	for l := 0; l < i; l++ {
		v := int(rand.Int63n(int64(i)))
		result = append(result, v)
	}

	sort.Ints(result)

	return result

}
