package _00_199

//189. 旋转数组
// TODO 第一种解决方案
/*
func rotate(nums []int, k int) {

    var first int
    var next int
    for i := 0; i < k; i++ {

        for i := 0; i < len(nums); i++ {

            if i+1 >= len(nums){
                first = 0
            }else {
                first = i +1
            }
            if i+2 >= len(nums){
                next = i+2 -len(nums)
            }else{
                next = i + 2
            }

            nums[i], nums[first] = nums[first], nums[next]
            fmt.Println(nums)

        }

    }
    fmt.Println(nums)

}
*/

//TODO 第二种解决方案 空间复杂度O(1)
func rotate(nums []int, k int) {

	var cnt int
	var lastval int
	for {

		if cnt >= len(nums){
			break
		}

		for i:=lastval;i<len(nums);i+=k{
			lastval = nums[lastval+k]


		}

		cnt+=len(nums)/k
		lastval++
	}

}

// TODO 第三种解决方案

func main() {
    //rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
