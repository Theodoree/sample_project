package main




var qwe []int

func inorderTraversal(root *TreeNode) []int {
	if root ==nil{
		return []int{}
	}

	LDR(root)
	ab := qwe

	return ab
}

func LDR(root *TreeNode) {
	if root != nil {
		LDR(root.Left) // L
		qwe = append(qwe,root.Val)
		LDR(root.Right) // R
	}
}

