package __99

import "fmt"

func permute(nums []int) [][]int {
    var result [][]int
    Permute(nums, []int{}, &result)
    return result
}

func Permute(nums, current []int, result *[][]int) []int {

    if len(nums) == 1 {
        return current
    }

    for i := 0; i < len(nums); i++ {

        current = append(current,Permute(append(nums[:i], nums[i+1:]...), append(current, nums[i]), result)...)
        *result = append(*result, current)

    }

    return  nil

}
func main() {
    fmt.Println(permute([]int{1, 2, 3}))
}
