package main

import (
    "fmt"
    "sort"
)

// 16.最接近的三数之和

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)

    var sum int
    var tmp int
    for i := 0; i < len(nums)-2; i++ {
        for j := i+1; j < len(nums)-1; j++ {
            tmp = nums[i] + nums[j] + nums[j+1]
            if tmp == target {
                return target
            }
            if tmp < target{
                sum = tmp
            }else{
                sumA:=tmp - target
                if sumA < 0 {
                    sumA = sumA*-1+1
                }
                sumB:=sum - target
                if sumB < 0 {
                    sumB = sumB*-1+1
                }

                if sumA < sumB{
                    return tmp
                } else{
                    return sum
                }
            }
        }
    }


    return tmp
}

func main() {

    f:=threeSumClosest([]int{-1,2,1,-4},1)
    fmt.Println(f)
    f = threeSumClosest([]int{1,1,1,1}, 0)
    fmt.Println(f)
}
