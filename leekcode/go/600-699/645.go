package _00_699



/*
645. 错误的集合

集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。

给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

示例 1:

输入: nums = [1,2,2,4]
输出: [2,3]
注意:

给定数组的长度范围是 [2, 10000]。
给定的数组是无序的。
*/

func findErrorNums(nums []int) []int {
    dup := 0
    for i:=0;i<len(nums);i++{
        n := abs(nums[i])

        if nums[n-1] < 0{
            dup = n
        }else{
            nums[n-1] = -nums[n-1]
        }
    }

    mis := 0
    for i, v := range nums{
        if v > 0{
            mis = i + 1
            break
        }
    }

    return []int{dup, mis}
}

func abs(num int) int {
    if num < 0 {
        return - num
    }
    return num
}