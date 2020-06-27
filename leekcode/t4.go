package main

import (
    "fmt"
    . "github.com/Theodoree/sample_project/leekcode/utils"
)

const (
    null = 0x7777777
)


func main() {
    t, _ := CreateTree([]int{3, 9, 20, null, null, 15, 7})
    fmt.Println(t)
    fmt.Println(maxSlidingWindow([]int{1},1))
}
