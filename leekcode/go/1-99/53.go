package __99


// 	53 最大子序和

func maxSubArray(nums []int) int {

	res := nums[0]
	var sum int

	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > res {
			res = sum
		}
	}

	return res
}

