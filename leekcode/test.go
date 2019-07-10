package main

import (
    "fmt"
)

/*
func pivotIndex(nums []int) int {

    switch {
    case len(nums) == 0:
        return -1
    case len(nums) == 1:
        return 0
    }

    var leftSum, left, rightSum int
    right := len(nums) - 1
    for left < right {
        if nums[left] == 0 {
            left++
            continue
        }
        if nums[right] == 0 {
            right--
            continue
        }
        if math.Abs(float64(rightSum)) > math.Abs(float64(leftSum)) {
            leftSum += nums[left]
            left++
        } else {
            rightSum += nums[right]
            right--
        }

    }

    if leftSum == rightSum {
        return left
    }

    return -1
}
*/

func pivotIndex(nums []int) int {

    var leftIndex, rightIndex int
    var left int
    var right int
    for k, _ := range nums {

        leftIndex = k - 1
        right = len(nums) - 1
        left = 0
        right = 0

        for leftIndex >= 0 {

            left += nums[leftIndex]
            leftIndex--
        }

        for rightIndex > k {
            right += nums[rightIndex]
            rightIndex--
        }
        if left == right {
            return k
        }
    }

    return -1

}
func main() {
    fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
    fmt.Println(pivotIndex([]int{-1, -1, 0, -1, -1, -1}))
    fmt.Println(pivotIndex([]int{-1, -1, -1, 0, -1, -1}))
    fmt.Println(pivotIndex([]int{-1, -1, -1, 0, 1, 1}))

}
