package __99

import "fmt"

//55 跳跃游戏
func canJump(nums []int) bool {
    if len(nums) == 0 {
        return true
    }
    if nums[0] == len(nums)-1 {
        return true
    }

    return Check(nums, 1, nums[0])

}

func Check(nums []int, n int, j int) bool {

    if len(nums) <= n+j {
        return true
    }

    for ; n < j+n; n++ {

        Check(nums, n, nums[n])

    }

    return false
}

func main() {
    fmt.Println(canJump([]int{2, 3, 1, 1, 4}))

}
