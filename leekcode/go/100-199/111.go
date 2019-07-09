package _00_199

import (
    "sort"
)

func minDepth(root *TreeNode) int {

    if root == nil {
        return 0
    }

    var s []int

    MinDepth(root, 0, &s)
    sort.Ints(s)
    if s[0] == 1 && len(s) > 1 {
        s[0] = s[1]
    }

    return s[0]
}

func MinDepth(root *TreeNode, depth int, s *[]int) {

    if root == nil {
        *s = append(*s, depth)
        return
    }

    MinDepth(root.Left, depth+1, s)
    MinDepth(root.Right, depth+1, s)

}