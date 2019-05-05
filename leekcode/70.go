package main

import "fmt"

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	first := 1
	second := 2

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

func climbStairs1(n int) int {
	if n == 1 || n == 2 || n == 0 {
		return n
	}
	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 2
	for i:=2; i<n; i++ {
		arr[i] = arr[i-1]+arr[i-2]
	}
	return arr[n-1]
}

func main() {
	fmt.Println(climbStairs(2))
}
