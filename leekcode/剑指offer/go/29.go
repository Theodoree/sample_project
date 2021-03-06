package _go

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
        return []int{}
    }

    var matrixBool [][]bool

    for i := 0; i < len(matrix); i++ {
        matrixBool = append(matrixBool, make([]bool, len(matrix[0])))
    }

    var left, right, up, down bool
    var row, col int
    var result []int
    right = true

    for len(result) < len(matrix)*len(matrix[0]) {

        switch {
        case left:
            if col < 0 || (col >= 0 && matrixBool[row][col]) {
                left = false
                up = true
                row--
                if col < 0 {
                    col = 0
                } else {
                    col++
                }
                continue
            } else {
                result = append(result, matrix[row][col])
                matrixBool[row][col] = true
                col--
            }

        case right:
            if col == len(matrix[0]) || (col < len(matrix[0]) && matrixBool[row][col]) {
                right = false
                down = true
                row++
                if col == len(matrix[0]) {
                    col--
                } else {
                    col--
                }
                continue
            } else {
                result = append(result, matrix[row][col])
                matrixBool[row][col] = true
                col++
            }
        case up:
            if row < 0 || (row >= 0 && matrixBool[row][col]) {
                up = false
                right = true
                col++
                if row < 0 {
                    row = 0
                } else {
                    row++
                }
                continue
            } else {
                result = append(result, matrix[row][col])
                matrixBool[row][col] = true
                row--
            }

        case down:
            if row == len(matrix) || (row < len(matrix) && matrixBool[row][col]) {
                down = false
                left = true
                col--
                row--
                continue
            } else {
                result = append(result, matrix[row][col])
                matrixBool[row][col] = true
                row++
            }
        }
    }

    return result
}
