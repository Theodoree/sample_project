package main

func twoSum(nums []int, target int) []int {
	i := make(map[int]int)
	for index, value := range nums {
		complement := target - nums[index]
		if value, ok := i[complement]; ok {
			return []int{value, index}
		}
		i[value] = index
	}
	return nil
}
