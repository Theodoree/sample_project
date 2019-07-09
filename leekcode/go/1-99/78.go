package __99

import "fmt"

//78.子集
func subsets(nums []int) [][]int {

    ans := make(map[string][]int)
    SubSets(nums, ans, []int{})
    fmt.Println(ans)
    return  nil
}

func SubSets(nums []int, ans map[string][]int, tem []int) {
    if nums == nil{
        return
    }

    for i := 0; i < len(nums); i++ {
        var key string
        tem = append(tem,nums[i])
        for _, t := range tem {
            key += fmt.Sprintf(" %d ", t)
        }
        fmt.Println(i)
        ans[key] = tem
        SubSets(nums[1:], ans, tem)
    }

}

func main(){
    subsets([]int{1,2,3})
}