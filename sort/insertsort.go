package main

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			j := i - 1
			temp := nums[i]
			for j >= 0 && nums[j] > temp {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] = temp
		}
	}
}