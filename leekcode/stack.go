package main

import (
    "fmt"
    
    . "github.com/Theodoree/sample_project/leekcode/utils"
)

func main() {
    
    v := ReverseList(CreateLinkNode([]int{1, 2, 3, 4}))
    
    for v != nil {
        fmt.Println(v.Val)
        v = v.Next
    }
    
}

func ReverseList(head *ListNode) *ListNode {
    //  前一个节点
    var preNode *ListNode
    preNode = nil
    //  后一个节点
    nextNode := &ListNode{}
    nextNode = nil
    for head != nil {
        // 1->2->3->4
        // nextNode 2->3->4
        // head 1->nil
        // preNode  1-> nil
        // head 2->3->4
        
        // nextNode 3->4
        // head 2->1->null
        // pre 2->1->null
        // head 3->4
        //  保存头节点的下一个节点，
        nextNode = head.Next
        //  将头节点指向前一个节点
        head.Next = preNode
        //  更新前一个节点
        preNode = head
        //  更新头节点
        head = nextNode
    }
    
    return preNode
}