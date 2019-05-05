package main

import "fmt"

//34. 在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {

	mid := len(nums) / 2
	for {

		if nums[mid] > target {
			mid = mid/2
		} else {
			break
		}
	}

	return findFirstLastIndex(nums, mid, target)

}

func findFirstLastIndex(nums []int, mid int, target int) []int {

	last := mid * 2
	for {

		if last <= mid {
			break
		}

		if nums[mid] != target {
			mid++
		}

		if nums[last] != target {
			last--
		}


		if nums[last] == target && nums[mid] == target {
			return  []int{mid,last}
		}
	}


	return  []int{-1,-1}
}
func main() {

	f := searchRange([]int{2, 2}, 2)
	fmt.Println(f)

}
