package _00_999


/*
961. 重复 N 次的元素

在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。

返回重复了 N 次的那个元素。
示例 1：

输入：[1,2,3,3]
输出：3
示例 2：

输入：[2,1,2,5,3,2]
输出：2
示例 3：

输入：[5,1,5,2,5,3,5,4]
输出：5
*/

func repeatedNTimes(A []int) int {
    //filterMap := make(map[int]int)
    //for _, v := range A {
    //    filterMap[v]++
    //    if filterMap[v] == len(A)/2 {
    //        return v
    //    }
    //}
    counting := [10001]int{}
    for _, val := range A {
        counting[val] =  counting[val] + 1
        if counting[val] == len(A)/2 {
            return val
        }
    }
    return 0
}
