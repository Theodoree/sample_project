package main

import (
    "encoding/json"
    "fmt"
    "github.com/piex/transcode"
)

const (
    null = 0x7777777
)

/* 自底而上
   1.DP方程 状态转移方程
   2.边界条件 初始状态
*/

/*
914. 卡牌分组

给定一副牌，每张牌上都写着一个整数。

此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：

每组都有 X 张牌。
组内所有的牌上都写着相同的整数。
仅当你可选的 X >= 2 时返回 true。



示例 1：

输入：[1,2,3,4,4,3,2,1]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
示例 2：

输入：[1,1,1,2,2,2,3,3]
输出：false
解释：没有满足要求的分组。
示例 3：

输入：[1]
输出：false
解释：没有满足要求的分组。
示例 4：

输入：[1,1]
输出：true
解释：可行的分组是 [1,1]
示例 5：

输入：[1,1,2,2,2,2]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[2,2]

提示：

1 <= deck.length <= 10000
0 <= deck[i] < 10000
*/

func hasGroupsSizeX(deck []int) bool {

    // 一个分组至少要有二个
    if len(deck) < 2 {
        return false
    }

    m := make(map[int]int)

    // 现将已有的数记入map
    for _, v := range deck {
        m[v]++
    }

    /*
       此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
       每组都有 X 张牌。
       组内所有的牌上都写着相同的整数。
       仅当你可选的 X >= 2 时返回 true。
    */

    var min int

    // 获取最小的元素的个数
    for _, v := range m {
        if min == 0 {
            min = v
        } else if v < min {
            min = v
        }
    }

    for min%2 == 0 && min/2 > 1 {
        min = min / 2
    }

    // 判断Map内的元素数量是否能整除min
    for _, v := range m {
        if v%min != 0 {
            return false
        }
    }

    return true
}



func main() {
    // val := []int{1,1,1,1,2,2,2,2,2,2}
    // fmt.Println(hasGroupsSizeX(val))


}
