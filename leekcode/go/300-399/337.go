package _00_399

import "sort"

func rob(root *TreeNode) int {

    var f []int
    backtrackNode(&f, 0, root)

    sort.Ints(f)
    return f[len(f)-1]
}

func backtrackNode(ret *[]int, sum int, node *TreeNode) {
    if node == nil {
        *ret = append(*ret, sum)
        return
    }

    backtrackNode(ret, sum+node.Val, node.Left)
    backtrackNode(ret, sum+node.Val, node.Right)

}
