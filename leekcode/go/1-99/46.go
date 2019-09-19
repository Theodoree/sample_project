package __99

import "fmt"

/*
46. 全排列
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
func permute(nums []int) [][]int {
    var result [][]int
    Permute(nums, &result, []int{})

    return result

}

func Permute(nums []int, result *[][]int, cur []int) {

    if len(nums) == 0 {
        *result = append(*result, cur)
        return
    }

    for i := 0; i < len(nums); i++ {

        c := make([]int, len(nums))
        copy(c, nums)
        c = append(c[:i], c[i+1:]...)
        Permute(c, result, append(cur, nums[i]))

    }

}


func main() {
    fmt.Println(permute([]int{1, 2, 3}))
}
