package _go


/*
面试题42. 连续子数组的最大和

输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

提示：

1 <= arr.length <= 10^5
-100 <= arr[i] <= 100
*/

func maxSubArray(nums []int) int {

    res := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > 0 {
            nums[i] += nums[i-1 ]
        }
        res = max(res,nums[i])

    }
    return  res
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

