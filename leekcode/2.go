package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var ListA = []int{}
var ListB = []int{}
var ListC = []int{}
var Head = &ListNode{}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	GetValue(l1,l2)
	fmt.Println(ListA,ListB)
	AddSum(ListA, ListB)
	fmt.Println(ListC)
	return GetNewListNode(ListC).Next

}

func GetValue(l1 *ListNode, l2 * ListNode) {
	for l1.Next !=nil{
		ListA = append(ListA,l1.Val)
		l1 = l1.Next
	}
	ListA = append(ListA,l1.Val)
	for l2.Next !=nil{
		ListB = append(ListB,l2.Val)
		l2 = l2.Next
	}
	ListB = append(ListB,l2.Val)
}

func GetNewListNode(l1 []int) *ListNode {
	now := Head
	for _, value := range l1 {
		node := ListNode{}
		node.Val = value
		node.Next = nil
		for now.Next != nil {
			now = now.Next
		}
		now.Next = &node
	}
	return Head
}

func AddSum(l1, l2 []int) {
	lenA := len(l1)
	lenB := len(l2)
	max := Max(lenA, lenB)
	var count int
	for i := 0; i < max; i++ {
		var valueA, valueB, sum int
		if i < lenA {
			valueA = l1[i]
		}
		if i < lenB {
			valueB = l2[i]
		}
		sum = valueA + valueB + count
		if sum >= 10 {
			ListC = append(ListC, sum%10)
			count = 1
		} else {
			ListC = append(ListC, sum)
			count = 0
		}
	}
}

func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func gouzao(i []int)*ListNode{
	Head:=&ListNode{}
	current:= Head
	for _,value:=range i{
		Node :=&ListNode{}
		Node.Val = value
		Node.Next = nil
		for current.Next !=nil{
			current = current.Next
		}
		current.Next = Node
	}
	return Head
	}
func main() {
	l1:=[]int{7,8,8,0,6,5,1,9,9,1}
	l2:=[]int{8,8,8,1,0,9,7,5,5}
	List1:=gouzao(l1)
	List2:=gouzao(l2)
	a:=addTwoNumbers(List1.Next,List2.Next)
	for a.Next !=nil{
		fmt.Println(a.Val)
		a = a.Next
	}
	fmt.Println(a.Val)
	}
