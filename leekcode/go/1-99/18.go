package __99

import "fmt"

func fourSum(nums []int, target int) [][]int {

    var result [][]int
    FourSum(nums, target, []int{}, &result)
    fmt.Println(result)
    return result
}

func FourSum(nums []int, target int, current []int, result *[][]int) {
    if len(current) == 4 {
        var sum int
        for _, v := range current {
            sum += v
        }
        if sum == target {
            *result = append(*result, current)
        }
        return
    }
    if len(current) > 4{
        return
    }

    for i := 0; i < len(nums); i++ {
        t:= append(current,nums[i])
        FourSum(nums[i:], target,t, result)
    }

}

func main() {
    fourSum([]int{1, 0, -1, 0, -2, 2}, 0)

}
