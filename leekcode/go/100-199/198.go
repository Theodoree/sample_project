package _00_199


/*
198. 打家劫舍
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。


回溯算法的本质是从上到下  遍历所有可能的解

动态划分是从下至上,最优解
*/


func rob(nums []int) int {

    switch len(nums) {
    case 2:
        if nums[0] > nums[1] {
            return nums[0]
        } else {
            return nums[1]
        }
    case 1:
        return nums[0]
    case 0:
        return 0

    }

    a := nums[0]
    b := nums[1]
    for i := 2; i < len(nums); i++ {
        tmp := a

        if b > a {
            a = b
        }

        b = tmp + nums[i]
    }

    if a > b {
        return a
    }
    return b

}
