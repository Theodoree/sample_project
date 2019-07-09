package _00_399

import "fmt"

func lengthOfLIS(nums []int) int {

    cnt := 1
    tem := 1
    for i:=1;i<len(nums);i++{

        if nums[i-1] < nums[i]{
            cnt++
        }else{
                if cnt > tem {
                    tem =cnt
                }
                cnt = 1
        }
    }

    return  tem

}



func main(){

    fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))

}