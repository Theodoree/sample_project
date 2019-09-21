package main

import (
    . "github.com/Theodoree/sample_project/leekcode/utils"
)

/*
255. 验证前序遍历序列二叉搜索树

给定一个整数数组，你需要验证它是否是一个二叉搜索树正确的先序遍历序列。

你可以假定该序列中的数都是不相同的。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [5,2,6,1,3]
输出: false
示例 2：

输入: [5,2,1,3,6]
输出: true
进阶挑战：

您能否使用恒定的空间复杂度来完成此题？
*/

type Stack struct {
    num  int
    Next *Stack
}

func verifyPreorder(preorder []int) bool {
    if len(preorder) == 0 {
        return false
    }

    var left = true
    top := Stack{num: preorder[0]}

    for i := 1; i < len(preorder); i++ {


    }

}

func main() {

    t, _ := CreateTree([]int{5, 2, 6, 1, 3})

    PrintTree(t, "DLR", nil)

}
