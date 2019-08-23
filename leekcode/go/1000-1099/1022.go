package _000_1099

/*
1022. 从根到叶的二进制数之和

给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。

对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。

以 10^9 + 7 为模，返回这些数字之和。



示例：



输入：[1,0,1,0,1,0,1]
输出：22
解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22


提示：

树中的结点数介于 1 和 1000 之间。
node.val 为 0 或 1 。

*/
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int { // 层序遍历

    if root == nil {
        return 0
    }

    return SumRootToLeaf(root, 0)
}

func SumRootToLeaf(root *TreeNode, num int) int {

    sum := 0
    num = (num << 1) + root.Val
    if root.Left == nil && root.Right == nil {
        return num
    }
    if root.Left != nil {
        sum += SumRootToLeaf(root.Left, num)
    }
    if root.Right != nil {
        sum += SumRootToLeaf(root.Right, num)
    }

    return sum
}
