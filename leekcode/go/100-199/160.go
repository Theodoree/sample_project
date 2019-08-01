package _00_199


/*
160. 相交链表

编写一个程序，找到两个单链表相交的起始节点。
*/

type ListNode struct {
    Val  int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

    filterMap := make(map[*ListNode]struct{})

    for headA != nil || headB != nil {

        if headA != nil {
            if _, ok := filterMap[headA]; ok {
                return headA
            } else {
                filterMap[headA] = struct{}{}
            }
            headA = headA.Next

        }
        if headB != nil {
            if _, ok := filterMap[headB]; ok {
                return headB
            } else {
                filterMap[headB] = struct{}{}
            }
            headB = headB.Next

        }
    }

    return nil

}