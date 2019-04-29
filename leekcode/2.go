package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	Head := &ListNode{}
	carry := 0
	current := Head
	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}

		sum := carry + val1 + val2
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return Head.Next
}