package main

import (
	. "github.com/Theodoree/sample_project/leekcode/utils"
)

const null = Null

func reverseBits(num int) int {

	if num == 0 {
		return 1
	}
	var lastZero = - 1
	var leftIndex = 0
	var rightIndex = 0
	max := 0
	cur := 0
	for i := 0; i <= 32; i++ {
		if 1<<i&num > 0 {
			leftIndex++
			cur++
		} else {
			if lastZero == -1 {
				lastZero = i
				if cur > 0 {
					rightIndex = cur
				}
				cur++
				continue
			}

			if cur > max {
				max = cur
			}

			lastZero = i
			cur, rightIndex = leftIndex+1-rightIndex, cur
			leftIndex = 0
		}
	}

	return max

}


func main() {
	
}
