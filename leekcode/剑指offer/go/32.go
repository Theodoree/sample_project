package _go


func levelOrder(root *TreeNode) [][]int {
    var result [][]int
    dfs(root, 0, &result)
    return result
}

func dfs(root *TreeNode, depth int, result *[][]int) {
    if root == nil {
        return
    }

    if len(*result) == depth {
        *result = append(*result, []int{root.Val})
    } else {
        v := *result
        v[depth] = append(v[depth], root.Val)
    }

    dfs(root.Left, depth+1, result)
    dfs(root.Right, depth+1, result)

}
