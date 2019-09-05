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

func lexicalOrder(n int) []int {

    result := make([]string, n, n)

    for i := 0; i < len(result); i++ {
        result[i] = strconv.Itoa(i + 1)
    }
    sort.Slice(result, func(i, j int) bool {
        return result[i] < result[j]
    })

    var re []int

    for _, v := range result {
        i, _ := strconv.Atoi(v)

        re = append(re, i)
    }
    return re
}

/*
328. 奇偶链表

给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/

func oddEvenList(head *ListNode) *ListNode {

    if head == nil || head.Next == nil || head.Next.Next == nil {
        return head
    }

    var j, o *ListNode
    var jcur, ocur *ListNode
    var cnt int
    cnt = 1
    for head != nil {

        if cnt%2 != 0 {
            if jcur == nil {
                j, jcur = head, head
            } else {
                jcur.Next = head
                jcur = jcur.Next
            }
        } else {
            if ocur == nil {
                o, ocur = head, head
            } else {
                ocur.Next = head
                ocur = ocur.Next
            }
        }
        head = head.Next
        cnt++
    }

    if jcur != nil {
        jcur.Next = o
        if ocur != nil {
            ocur.Next = nil
        }
    }
    fmt.Println(j)
    return head
}

func main() {

    n1 := 0b00101101
    fmt.Println(n1)

    /*

       n1 := &ListNode{Val: 1}
       n2 := &ListNode{Val: 2}
       n3 := &ListNode{Val: 3}
       n4 := &ListNode{Val: 4}
       n5 := &ListNode{Val: 5}
       n1.Next = n2
       n2.Next = n3
       n3.Next = n4
       n4.Next = n5

       n1 = oddEvenList(n1)
       for n1 != nil {
           fmt.Println(n1.Val)
           n1 = n1.Next
       }
    */
}

func readXlsx() {
    xlsx, err := excelize.OpenFile("/Users/ted/Downloads/京东联盟部分高佣商品列表副本.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    cell := xlsx.GetCellValue("Sheet1", "B2")
    fmt.Println(cell)

    var result []interface{}
    rows := xlsx.GetRows("Sheet1")
    for i := 1; i < len(rows); i++ {
        row := rows[i]
        price := strings.Split(row[2], `￥`)
        if len(price) > 1 {
            row[2] = price[1]
        } else {
            row[2] = price[0]
        }

        Sku := struct {
            Title    string
            OriPrice string
            Price    string
            UnionUrl string
            Img      string
        }{
            Title:    row[1],
            OriPrice: row[2],
            Price:    row[2],
            UnionUrl: row[4],
            Img:      row[3],
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
