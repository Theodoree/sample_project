package main

//189. 旋转数组
// TODO 第一种解决方案
func rotate(nums []int, k int) {

	var index int
	array := make([]int, len(nums))
	for key, val := range nums {

		if key+k >= len(nums) {
			index = key + k - (len(nums))
		} else {
			index = key + k
		}

		array[index] = val
	}

	nums = array
}

// TODO 第二种解决方案 空间复杂度O(1)
//func rotate(nums []int, k int) {
//
//	var cnt int
//	var lastval int
//	for {
//
//		if cnt == len(nums){
//			break
//		}
//
//		for i:=lastval;i<len(nums);i+=k{
//			lastval = nums[lastval+k]
//
//
//		}
//
//		cnt+=len(nums)/k
//		lastval++
//	}
//
//}

// TODO 第三种解决方案

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)

}
