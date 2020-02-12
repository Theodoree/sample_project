package main

import (
    "fmt"
    . "github.com/Theodoree/sample_project/leekcode/utils"
)

const (
    null = 0x7777777
)

/*
面试题29. 顺时针打印矩阵

输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [
[1,2,3],
[4,5,6],
[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100
*/

func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil
    }
    var result []int
    var top, down, left, right bool
    var minCol, maxCol, minIndex, maxIndex int

    maxCol = len(matrix) - 1
    maxIndex = len(matrix[0]) - 1
    var col, index int
    right = true
    result = append(result, matrix[0][0])
    for len(result) < len(matrix)*len(matrix[0]) {
        switch {
        case top:
            if col == minCol {
                right = true
                top = false
                minCol++
                continue
            }
            col--
            result = append(result, matrix[col][index])
        case down:
            if col == maxCol {
                down = false
                left = true
                maxCol--
                continue
            }
            col++
            result = append(result, matrix[col][index])
        case left:
            if index == minIndex {
                top = true
                left = false
                minIndex++
                continue
            }
            index--
            result = append(result, matrix[col][index])
        case right:
            if index == maxIndex {
                down = true
                right = false
                maxIndex--
                minCol++
                continue
            }
            index++
            result = append(result, matrix[col][index])
        }


    }

    return result

}
func main() {
    t := CreateLinkNode([]int{1, 2, 4})
    t1 := CreateLinkNode([]int{1, 3, 4})

    fmt.Println(t, t1)

    var f = [][]int{
        {1,2,3,4,5},
        {6,7,8,9,10},
        {11,12,13,14,15},
        {16,17,18,19,20},
        {21,22,23,24,25}}
    //f = [][]int{{3}, {2}}
    fmt.Println(spiralOrder(f))
    //1,2,3,4,5,10,15,20,25,24,23,22,21,16,11,6,7,8,9,14,19,18,17,12,13
    //1 2 3 4 5 10 15 20 25 24 23 22 21 16 11 6 7 8 9 14 19 18 17 18 23
}
