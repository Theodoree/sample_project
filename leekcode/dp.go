package main

import (
    "fmt"

    . "github.com/Theodoree/sample_project/leekcode/utils"
)

const (
    null = 0x7777777
)

/* 自底而上
   1.DP方程 状态转移方程
   2.边界条件 初始状态
*/


func main() {

    m := CreateLinkNode([]int{9})

    for m != nil {
        fmt.Print(m.Val, "->")
        m = m.Next
    }
}
