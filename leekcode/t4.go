package main

import (
    "encoding/json"
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
    "math/rand"
    "sort"
    "strings"
    "time"
)

type Sku struct {
    Title    string `json:"title"`
    OriPrice string `json:"ori_price"`
    Price    string `json:"price"`
    UnionUrl string `json:"union_url"`
    Img      string `json:"img"`
}

func main() {



}

func readXlsx() {
    xlsx, err := excelize.OpenFile("/Users/ted/Downloads/京东联盟部分高佣商品列表副本.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    cell := xlsx.GetCellValue("Sheet1", "B2")
    fmt.Println(cell)

    var result []*Sku
    rows := xlsx.GetRows("Sheet1")
    for i := 1; i < len(rows); i++ {
        row := rows[i]
        price := strings.Split(row[2], `￥`)
        if len(price) > 1 {
            row[2] = price[1]
        } else {
            row[2] = price[0]
        }

        result = append(result, &Sku{
            Title:    row[1],
            OriPrice: row[2],
            Price:    row[2],
            UnionUrl: row[4],
            Img:      row[3],
        })
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
