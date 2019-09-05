package _00_399

import (
    "math/rand"
)

/*
382. 链表随机节点

给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。

进阶:
如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？

示例:

// 初始化一个单链表 [1,2,3].
ListNode head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
Solution solution = new Solution(head);

// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
solution.getRandom();


蓄水池抽样
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    Head *ListNode
}

type ListNode struct {
    Next *ListNode
    Val  int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    return Solution{Head: head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {

    count, res := 1, -1
    cur := this.Head
    for cur != nil {
        if int(rand.Int63n(int64(count))) == 0 {
            res = cur.Val
        }
        count++
        cur = cur.Next
    }

    return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
