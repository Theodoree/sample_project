package _00_499

import "fmt"

func thirdMax(nums []int) int {

    result := make([]int, 4, 4)
    for i := 0; i < len(nums); i++ {
        result[0] = nums[i]
        if result[0] != result[2] && result[0] != result[3] && result[0] != result[1] {

            if (result[3] != 0 && result[2] != 0) || result[1] == 0 {
                result[1] = nums[i]
            }
        }
        if result[2] < result [1] {
            result[2], result[1] = result[1], result[2]
        }
        if result[3] < result[2] {
            result[3], result[2] = result[2], result[3]
        }
        fmt.Println(result)
    }
    if result[1] == 0 {
        return result[3]
    }

    return result[1]
}

func main() {
    fmt.Println(thirdMax([]int{-2147483648,1,1}))

}
