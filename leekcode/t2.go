package main

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


func main() {

}
