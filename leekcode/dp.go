package main

import (
    "fmt"
    . "github.com/Theodoree/sample_project/leekcode/utils"
    "math/rand"
    "time"
)

type Solution struct {
    curNode  *ListNode
    headNode *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    return Solution{
        curNode:  head,
        headNode: head,
    }

}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    if this.curNode == nil {
        this.curNode = this.headNode
    }

    val := this.curNode.Val
    this.curNode = this.curNode.Next
    return val
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {

    n := CreateLinkNode([]int{0, 1, 2, 3, 4})
    fmt.Println(n == nil)

    for i := 0; i < 10; i++ {

        fmt.Println("刷屏金箍棒")
    }

}
