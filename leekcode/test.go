package main

func findRepeatNumber(nums []int) int {

    var arr  = make([]int,len(nums)+1)
    for _,v:=range nums{
        arr[v]++

        if arr[v] >= 2 {
            return  v
        }
    }

    for k,v:=range arr{
        if v >= 2 {
            return  k
        }
    }
    return  0
}

func main() {

}
