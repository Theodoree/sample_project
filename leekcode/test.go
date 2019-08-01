package main

import (
    "fmt"
    "strconv"
)

func findDiagonalOrder(matrix [][]int) []int {

    if len(matrix) == 0 {
        return nil
    }

    var result []int
    if len(matrix[0]) == 1 || len(matrix) == 1 || (len(matrix) == 2 && len(matrix[0]) == 2) {
        for _, v := range matrix {
            result = append(result, v...)
        }
        return result
    }
    var flag bool
    m := len(matrix)
    result = append(result, matrix[0][0])
    for i := 1; i <= m; i++ {
        var l, h int

        if !flag {
            l = i
            h = 0

            if i == m {
                h = i - 1
            }

            for l > len(matrix[0]) {
                l--
            }

            for l >= 0 && h < m && h >= 0 {
                result = append(result, matrix[h][l])
                h++
                l--
            }
            flag = !flag

        } else {
            h = i
            l = 0

            if i == m {
                //l = i - 2
                h = i - 1
            }

            for l < len(matrix[0]) && h >= 0 {
                result = append(result, matrix[h][l])
                h--
                l++
            }
            flag = !flag
            fmt.Println(result)
        }
    }
    if len(matrix[0]) != 1 {
        result = append(result, matrix[len(matrix)-1][len(matrix[0])-1])
    }
    return result

}

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
func isValidSudoku(board [][]byte) bool {

    filterMap := make(map[int]int)

    for i := 0; i < len(board); i++ {
        filterMap = make(map[int]int)
        //数字 1-9 在每一行只能出现一次。
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] != '.' {
                num, _ := strconv.Atoi(string(board[i][j]))
                if _, ok := filterMap[num]; ok {
                    return false
                }
                filterMap[num]++
            }
        }

        filterMap = make(map[int]int)
        //数字 1-9 在每一列只能出现一次。
        for j := 0; j < len(board); j++ {
            if board[i][j] != '.' {
                num, _ := strconv.Atoi(string(board[j][i]))
                if _, ok := filterMap[num]; ok {
                    return false
                }
                filterMap[num]++
            }
        }

    }
    //数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
    filterMap = make(map[int]int)
    //var slice []int
    for i := 0; i < len(board); i = i + 3 {

        for j := i; j < i+3; j++ {
            //slice = append(slice, board[0][0])
        }
    }

    fmt.Printf("%s", board)
    return false
}

/*
给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。

数学表达式如下:

如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。
*/
func increasingTriplet(nums []int) bool {

    if len(nums) < 3 {
        return false
    }

    var first, second, third int

    for i := 0; i < len(nums); i++ {
        if nums[i] != third && nums[i] != second && nums[i] != first {
            third = nums[i]
            if third > second {
                second, third = third, second
                if second != 0 {
                    second = i
                }
            }

            if second > first {
                first, second = second, first
            }
        }
    }
    if first > second && second > third {
        return true
    }
    return false
}

/* 你想要的O(1) */

/*
var filterMap = map[byte]int{
   'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5,
   'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10,
   'K': 11, 'L': 12, 'M': 13, 'N': 14, 'O': 15,
   'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20,
   'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25,
   'Z': 26,
}

func titleToNumber(s string) int {

    var sum int

    if len(s) > 1 {
        sum = (26 * 26)

        for i := 1; i < len(s)-1; i++ {
            sum += filterMap[s[i]] * 26
        }
        sum += filterMap[s[len(s)-1]]
    } else {
        sum = filterMap[s[0]]
    }
    return sum
}
*/

func fractionToDecimal(numerator int, denominator int) string {

    n := float64(numerator)
    d := float64(denominator)
    var result string
    if n/d > 0 {

        result = fmt.Sprintf("%.2f", n/d)

    }
    return result

}

func main() {
    //703
    //fmt.Println(titleToNumber(`AAA`))
    //fmt.Println(titleToNumber(`ZY`))

    //fmt.Println(increasingTriplet([]int{5, 1, 5, 5, 2, 5, 4}))
    //fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))

    board := [][]byte{
        {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
    isValidSudoku(board)

    //fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
    //fmt.Println(findDiagonalOrder([][]int{{1}}))
    //fmt.Println(findDiagonalOrder([][]int{{3}, {2}}))
    //fmt.Println(findDiagonalOrder([][]int{{7}, {9}, {6}}))
    //fmt.Println(findDiagonalOrder([][]int{{6, 9, 7}}))
    //fmt.Println(findDiagonalOrder([][]int{{1, 2}, {3, 4,}}))
    //fmt.Println(findDiagonalOrder([][]int{{2, 5, 8}, {4, 0, -1}}))

}
