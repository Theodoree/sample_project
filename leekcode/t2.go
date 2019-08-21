package main

import "fmt"

/*
188. 买卖股票的最佳时机 IV

给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [2,4,1], k = 2
输出: 2
解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2:

输入: [3,2,6,5,0,3], k = 2
输出: 7
解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
*/

func maxProfit(k int, prices []int) int {

    if k > len(prices) {
        return 0
    }

    dp := make([][]int, len(prices)-1)

    for i := 0; i < len(prices)-1; i++ {
        result := make([]int, k)
        result[0] = prices[i]
        for j := 1; j < k && j+i < len(prices); j++ {
            result[j] = prices[j+i] - prices[j+i-1]
        }
        dp[i] = result
    }

    var max int
    maxIndex := len(dp[0]) - 1
    for i := 0; i < len(dp)-1; i++ {
        if dp[i][1] > 0 || dp[i+1][1] > 0 {
            if dp[i][maxIndex] > dp[i+1][maxIndex] {
                max += dp[i][1]
            } else {
                max += dp[i+1][1]
            }
            i++
        }
    }

    return max
}

/*
554. 砖墙

提交记录
你的面前有一堵方形的、由多行砖块组成的砖墙。 这些砖块高度相同但是宽度不同。你现在要画一条自顶向下的、穿过最少砖块的垂线。

砖墙由行的列表表示。 每一行都是一个代表从左至右每块砖的宽度的整数列表。

如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你需要找出怎样画才能使这条线穿过的砖块数量最少，并且返回穿过的砖块数量。

你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。
*/

func leastBricks(wall [][]int) int {

    return 0
}

func fourSum(nums []int, target int) [][]int {

    if len(nums) < 4 {
        return nil
    }
    var result [][]int

    return result
}

/*
1041. 困于环中的机器人

在无限的平面上，机器人最初位于 (0, 0) 处，面朝北方。机器人可以接受下列三条指令之一：

"G"：直走 1 个单位
"L"：左转 90 度
"R"：右转 90 度
机器人按顺序执行指令 instructions，并一直重复它们。

只有在平面中存在环使得机器人永远无法离开时，返回 true。否则，返回 false。
*/
func isRobotBounded(instructions string) bool {

    return false
}

/*
696. 计数二进制子串

给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。

示例 1 :

输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

请注意，一些重复出现的子串要计算它们出现的次数。

另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
示例 2 :

输入: "10101"
输出: 4
解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
*/

func countBinarySubstrings(s string) int {
    return 0
}

/*
405. 数字转换为十六进制数

给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 补码运算 方法。

注意:

十六进制中所有字母(a-f)都必须是小写。
十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
给定的数确保在32位有符号整数范围内。
不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。
示例 1：

输入:
26

输出:
"1a"
示例 2：

输入:
-1

输出:
"ffffffff"
*/

func toHex(num int) string {

    return ""
}



func isCousins(root *TreeNode, x int, y int) bool {
    var result [][]int
    IsCousins(root, root, 1, x, y, &result)
    var left, right bool
    if len(result) > 0 {

        for _, v := range result[len(result)-1] {

            if v == x {
                left = true
            }

            if v == y {
                right = true
            }

            if left == true && right == true {
                return true
            }

        }
    }

    return left == true && right == true

}

func IsCousins(root *TreeNode, Parent *TreeNode, depth, x, y int, result *[][]int) {

    if root == nil {
        return
    }

    if len(*result) <= depth-1 {
        *result = append(*result, []int{})
    }
    current := *result

    if Parent.Left != nil && Parent.Right != nil {
        if !((Parent.Left.Val == x && Parent.Right.Val == y) || (Parent.Right.Val == x && Parent.Left.Val == y)) {
            current[depth-1] = append(current[depth-1], root.Val)
        }
    } else {
        current[depth-1] = append(current[depth-1], root.Val)
    }

    if root.Val == x || root.Val == y {
        return
    }

    IsCousins(root.Left, root, depth+1, x, y, result)
    IsCousins(root.Right, root, depth+1, x, y, result)

}

func main() {

    r1 := &TreeNode{Val: 1}
    r2 := &TreeNode{Val: 2}
    r3 := &TreeNode{Val: 3}
    r4 := &TreeNode{Val: 4}
    r5 := &TreeNode{Val: 5}

    r1.Left = r2
    r2.Right = r4
    r1.Right = r3
    r3.Right = r5
    fmt.Println(isCousins(r1, 5, 4))
}
