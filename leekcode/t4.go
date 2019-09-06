package main

import (
    "encoding/json"
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
    "math/rand"
    "sort"
    "strconv"
    "strings"
    "time"
)




func main() {
    v := [][]int{{1, 2, 5},
        {3, 2, 1}}

    fmt.Println(minPathSum(v))
}

func readXlsx() {
    xlsx, err := excelize.OpenFile("/Users/ted/Downloads/兰芝-微信直播商品-汇总.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }

    var result []interface{}
    rows := xlsx.GetRows("直播间选品")

    fmt.Println(rows[0])
    var unionIndex, titleIndex, priceIndex, imgIndex int

    for k, v := range rows[0] {
        switch v {
        // case "sku_id":
        //     skuIndex = k
        case "title":
            titleIndex = k
        case "price":
            priceIndex = k
        case "img":
            imgIndex = k
        case "union_url":
            unionIndex = k

        }
    }

    for i := 1; i < len(rows); i++ {
        row := rows[i]
        price := strings.Split(row[priceIndex], `￥`)
        if len(price) > 1 {
            row[priceIndex] = price[1]
        } else {
            row[priceIndex] = price[0]
        }

        Sku := struct {
            Title    string
            OriPrice string
            Price    string
            UnionUrl string
            Img      string
        }{
            Title:    row[titleIndex],
            OriPrice: row[priceIndex],
            Price:    row[priceIndex],
            UnionUrl: row[unionIndex],
            Img:      row[imgIndex],
        }
        result = append(result, &Sku)
    }

    b, _ := json.Marshal(result)
    fmt.Printf("%s", b)

}

const (
    DLR   = "DLR"
    LDR   = "LDR"
    LRD   = "LRD"
    LEVEL = "LEVEL"
    null  = 0x7777777
)

type ListNode struct {
    Val  int
    Next *ListNode
}
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func CreateTree(b []int) (*TreeNode, []*TreeNode) {

    var root *TreeNode
    if len(b) == 0 {
        return nil, nil
    }
    root = &TreeNode{Val: b[0]}

    v := make([]*TreeNode, len(b)+1)
    var index, index1 int
    index = 1
    index1 = 1
    current := root
    v[0] = root
    for i := 1; i <= len(b)/2; i++ {

        if b[i*2-1] != null {
            current.Left = &TreeNode{Val: b[i*2-1]}
            v[index1] = current.Left
            index1++
        }
        if i*2 < len(b) && b[i*2] != null {
            current.Right = &TreeNode{Val: b[i*2]}
            v[index1] = current.Right
            index1++
        }
        current = v[index]
        index++
    }

    return root, v

}
func LevelTree(root *TreeNode, result *[][]int, depth int) {

    if root == nil {
        return
    }

    if len(*result) <= depth {
        *result = append(*result, []int{})
    }

    k := *result
    k[depth] = append(k[depth], root.Val)
    LevelTree(root.Left, result, depth+1)
    LevelTree(root.Right, result, depth+1)

}

func PrintTree(root *TreeNode, Type string, array *[]int) {
    if root == nil {
        return
    }

    switch Type {
    case "DLR":
        if array != nil {
            *array = append(*array, root.Val)
        } else {
            fmt.Printf(" %d ", root.Val)
        }
        PrintTree(root.Left, Type, array)
        PrintTree(root.Right, Type, array)
    case "LDR":
        PrintTree(root.Left, Type, array)
        if array != nil {
            *array = append(*array, root.Val)
        } else {
            fmt.Printf(" %d ", root.Val)
        }
        PrintTree(root.Right, Type, array)
    case "LRD":
        PrintTree(root.Left, Type, array)
        PrintTree(root.Right, Type, array, )
        if array != nil {
            *array = append(*array, root.Val)
        } else {
            fmt.Printf(" %d ", root.Val)
        }
    }
}

// 二分查找，O(log(n / 2 + 1))
func arrangeCoins3(n int) int {
    l, r := 0, n/2+1
    for r-l > 1 {
        mid := (l + r) >> 1
        if mid*(mid+1) == 2*n {
            return mid

        } else if mid*(mid+1) <= 2*n {
            l = mid
        } else {
            r = mid
        }
    }

    if r*(r+1) == 2*n {
        return r
    }

    return l
}

func GetSlice(i int) []int {
    rand.Seed(time.Now().Unix())

    var result []int
    for l := 0; l < i; l++ {
        v := int(rand.Int63n(int64(i)))
        result = append(result, v)
    }

    sort.Ints(result)

    return result

}


/*
64. 最小路径和

给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/

func minPathSum(grid [][]int) int {

    if len(grid) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    dp := make([]int, (m+n)*2)

    dp[0] = grid[0][0]

    for i:=0;i<len(grid);i++{


    }

    return 1000

}

func MinPathSum(grid [][]int, m, n, current int, val string, min *int) {
    if m == len(grid)-1 && n == len(grid[0])-1 {
        if current < *min {
            *min = current
        }
        return
    }

    if m < len(grid)-1 {
        MinPathSum(grid, m+1, n, current+grid[m+1][n], val+" "+strconv.Itoa(grid[m+1][n]), min)
    }

    if n < len(grid[0])-1 {
        MinPathSum(grid, m, n+1, current+grid[m][n+1], val+" "+strconv.Itoa(grid[m][n+1]), min)
    }

}
