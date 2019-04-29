package main

import (
	"fmt"
)

//三数之和

func threeSum(nums []int) [][]int {

	//if len(nums) < 3 {
	//	return [][]int{}
	//}
	//
	//sort.Ints(nums)
	//
	//var array [][]int
	//for i := 0; i < len(nums); i++ {
	//	if i == 0 || nums[i] > nums[i-1]{
	//		left := 0
	//		right := len(nums) - 1
	//		for left < right && right < len(nums) {
	//			s := nums[left] + nums[right] + nums[i]
	//			if s == 0 {
	//				if i != left && left != right {
	//					array = append(array, []int{nums[i], nums[left], nums[right]})
	//				}
	//				left++
	//				right--
	//				for left < right && nums[left] == nums[left-1]{
	//					left++
	//				}
	//				for right > left && nums[right] == nums[right+1]{
	//					right--
	//				}
	//			} else if s > 0 {
	//				right--
	//			} else {
	//				left++
	//			}
	//
	//		}
	//	}
	//
	//
	//}
	//
	//return array
	//TODO 待写
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

}
