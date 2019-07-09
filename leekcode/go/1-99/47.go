package __99

import "fmt"

//47. 全排列 II

func permuteUnique(nums []int) [][]int {
    var result [][]int

    PermuteUnique(nums, []int{}, 0, len(nums), &result)
    return result
}

func PermuteUnique(nums []int, current []int, left, right int, result *[][]int) {
    if left == right {
        *result = append(*result, nums)
        return
    }

    for i := left; i < right; i++ {
        if i != left && nums[left] == nums[i] {
            continue
        }

        nums[left], nums[i] = nums[i], nums[left]
        PermuteUnique(nums, left+1, right, result)
    }

}

func main() {
    fmt.Println(permuteUnique([]int{1, 1, 2}))

}
