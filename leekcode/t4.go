package main

import (
    "encoding/json"
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
    "io/ioutil"
    "math/rand"
    "sort"
    "strconv"
    "strings"
    "time"

    . "github.com/Theodoree/sample_project/leekcode/utils"
)

const (
    null = 0x7777777
)

func main() {

    /*
       excelPath := "/Users/ted/Desktop/尚德8月消耗成本.xlsx"
       WritePath := "/Users/ted/Downloads/ted"
       readXlsxBy(excelPath, WritePath)
    */

    strs := strings.Split(funs, `,`)
    vals := strings.Split(val, `],`)
    const objName=`linkedList`
    for i := 0; i < len(strs); i++ {

        funcName := strings.ToUpper(string(strs[i][1])) + strs[i][2:len(strs[i])-1]
        fmt.Printf("%s.%s(%s) \n",objName,funcName,vals[i][1:])
    }

    fmt.Println(len(strs), len(vals))

}

const funs = `"addAtHead","get","addAtTail","deleteAtIndex","addAtHead","deleteAtIndex","get","addAtTail","addAtHead","addAtTail","addAtTail","addAtTail","addAtIndex","get","addAtIndex","addAtHead","deleteAtIndex","addAtIndex","addAtHead","addAtIndex","deleteAtIndex","get","addAtTail","deleteAtIndex","deleteAtIndex","addAtTail","addAtTail","addAtIndex","addAtHead","get","get","addAtTail","addAtTail","addAtTail","addAtTail","addAtIndex","addAtIndex","addAtHead","addAtIndex","addAtTail","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","addAtHead","deleteAtIndex","addAtHead","get","addAtHead","get","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtTail","deleteAtIndex","get","addAtIndex","addAtHead","addAtTail","deleteAtIndex","addAtHead","addAtIndex","deleteAtIndex","deleteAtIndex","deleteAtIndex","addAtHead","addAtTail","addAtTail","addAtHead","addAtTail","addAtIndex","deleteAtIndex","deleteAtIndex","addAtIndex","addAtHead","addAtHead","addAtTail","get","addAtIndex","get","addAtHead","addAtHead","addAtHead","addAtIndex","addAtIndex","get","addAtHead","get","get","addAtTail","addAtHead","addAtHead","addAtTail","addAtTail","get","addAtTail"`
const val = `[8],[1],[81],[2],[26],[2],[1],[24],[15],[0],[13],[1],[6,33],[6],[2,91],[82],[6],[4,11],[3],[7,14],[1],[6],[99],[11],[7],[5],[92],[7,92],[57],[2],[6],[39],[51],[3],[22],[5,26],[9,52],[69],[5,58],[79],[7],[41],[33],[88],[44],[8],[72],[93],[18],[1],[9],[46],[9],[92],[71],[69],[11,54],[27],[83],[12],[20],[19,97],[77],[36],[3],[35],[16,68],[22],[36],[17],[62],[89],[61],[6],[92],[28,69],[23],[28],[7,4],[0],[24],[52],[1],[23,3],[7],[6],[68],[79],[45,90],[41,52],[28],[25],[9],[32],[11],[90],[24],[98],[36],[34],[26],`

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

type Stats struct {
    Date  string `json:"date,omitempty"`
    Click uint64 `json:"click,omitempty"`
    Pv    uint64 `json:"pv,omitempty"`

    CTR string  `json:"ctr,omitempty"`
    CPM string  `json:"cpm,omitempty"`
    CPC float64 `json:"cpc,omitempty"`

    Cost    float64   `json:"cost,omitempty"`
    Balance float64   `json:"balance,omitempty"`
    Deposit []float64 `json:"deposit,omitempty"`
}

func readXlsxBy(ExcelPath, WritePath string) {
    xlsx, err := excelize.OpenFile(ExcelPath)
    if err != nil {
        fmt.Println(err)

    }

    // 读取的表名必须为sheet
    rows := xlsx.GetRows(`sheet`)
    if len(rows) == 0 {
        return
    }

    recordMap := make(map[string]*Stats)

    buf, err := ioutil.ReadFile(WritePath)
    if err == nil {
        var arr []*Stats
        err = json.Unmarshal(buf, &arr)
        for _, v := range arr {
            recordMap[v.Date] = v
        }

    }
    var filterRows [][]string
    for _, v := range rows {
        // 这里查看有没有历史数据
        if _, ok := recordMap[convertToFormatDay(v[0])]; !ok {
            filterRows = append(filterRows, v)
        }
    }
    rows = filterRows

    var (
        depositIndex = -1
        ClickIndex   = -1
        CPCIndex     = -1
        costIndex    = -1
        balanceIndex = -1
    )
    /*
       [{"date":"2019-08-28","click":115060,"pv":6374515,"ctr":"1.805","cpm":"10.830","cpc":0.6,"cost":69036,"balance":230964,"deposit":[100000,200000]},{"date":"2019-08-29","click":125841,"pv":7612885,"ctr":"1.653","cpm":"9.918","cpc":0.6,"cost":75504.6,"balance":155459.4},{"date":"2019-08-30","click":121281,"pv":5699295,"ctr":"2.128","cpm":"12.768","cpc":0.6,"cost":72768.6,"balance":82690.8},{"date":"2019-08-31","click":137514,"pv":6348753,"ctr":"2.166","cpm":"12.996","cpc":0.6,"cost":82508.4,"balance":182.400000000038}]
       [{"date":"2019-08-31","click":137514,"pv":6348753,"ctr":"2.166","cpm":"12.996","cpc":0.6,"cost":82508.4,"balance":182.400000000038},{"date":"2019-08-28","click":115060,"pv":6374515,"ctr":"1.805","cpm":"10.830","cpc":0.6,"cost":69036,"balance":230964,"deposit":[100000,200000]},{"date":"2019-08-29","click":125841,"pv":7612885,"ctr":"1.653","cpm":"9.918","cpc":0.6,"cost":75504.6,"balance":155459.4},{"date":"2019-08-30","click":121281,"pv":5699295,"ctr":"2.128","cpm":"12.768","cpc":0.6,"cost":72768.6,"balance":82690.8}]
    */
    rand.Seed(time.Now().UnixNano())
    for _, v := range rows {
        if v[0] == "" {
            continue
        }

        if v[0] == "日期" {
            for i, val := range v {
                switch val {
                case "充值":
                    depositIndex = i
                case "点击量":
                    ClickIndex = i
                    depositIndex = -1
                case "CPC":
                    CPCIndex = i
                case "消耗":
                    costIndex = i
                case "余额":
                    balanceIndex = i
                }
            }

        } else {
            record := &Stats{}

            v[0] = convertToFormatDay(v[0])

            if val, ok := recordMap[v[0]]; ok {
                record = val
            }

            record.Date = v[0]
            if depositIndex >= 0 {
                deposit, err := strconv.ParseFloat(v[depositIndex], 64)
                if err != nil {
                    fmt.Println("depositIndex", err)
                }
                record.Deposit = append(record.Deposit, deposit)
            }

            if ClickIndex >= 0 {
                click, err := strconv.ParseUint(v[ClickIndex], 10, 64)
                if err != nil {
                    fmt.Println(err)
                }
                record.Click = click
            }

            if CPCIndex >= 0 {
                CPC, err := strconv.ParseFloat(v[CPCIndex], 64)
                if err != nil {
                    fmt.Println("CPCIndex", err)
                }
                record.CPC = CPC
            }

            if costIndex >= 0 {
                cost, err := strconv.ParseFloat(v[costIndex], 64)
                if err != nil {
                    fmt.Println("costIndex", err)
                }
                record.Cost = cost
            }

            if balanceIndex >= 0 {
                balance, err := strconv.ParseFloat(v[balanceIndex], 64)
                if err != nil {
                    fmt.Println("balanceIndex", err)
                }
                record.Balance = balance
            }
            if record.CPC > 0 && record.Click > 0 && record.Cost > 0 {
                var rate = 1.9
                f := rand.Intn(30)
                rate *= (float64(f-15) / 100) + 1
                record.CTR = fmt.Sprintf("%.3f", rate)
                record.Pv = uint64(float64(record.Click) / (rate / 100))
                record.CPM = fmt.Sprintf("%.3f", record.Cost/float64(record.Pv)*1000)
            }

            recordMap[v[0]] = record
        }

    }

    var recordArr []*Stats

    for _, v := range recordMap {
        recordArr = append(recordArr, v)
    }

    buf, err = json.Marshal(&recordArr)
    if err != nil {
        fmt.Println(err)
    }

    err = ioutil.WriteFile(WritePath, buf, 0644)
    fmt.Println(err)

}

func convertToFormatDay(excelDaysString string) string {
    baseDiffDay := 38719
    curDiffDay := excelDaysString
    b, _ := strconv.Atoi(curDiffDay)
    realDiffDay := b - baseDiffDay
    realDiffSecond := realDiffDay * 24 * 3600
    baseOriginSecond := 1136185445
    resultTime := time.Unix(int64(baseOriginSecond+realDiffSecond), 0).Format("2006-01-02")
    return resultTime
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

    for i := 0; i < len(grid); i++ {

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

/*
   86. 分隔链表

   给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

   你应当保留两个分区中每个节点的初始相对位置。

   示例:

   输入: head = 1->4->3->2->5->2, x = 3
   输出: 1->2->2->4->3->5
*/

func partition(head *ListNode, x int) *ListNode {

    if head == nil {
        return head
    }

    cur := head
    min := &ListNode{Val: 0}
    minCur := min
    for cur.Next != nil {

        if cur.Val == x {
            for cur.Next != nil {
                if cur.Next.Val < x {
                    minCur.Next = cur.Next
                    minCur = minCur.Next
                    cur.Next = cur.Next.Next
                } else {
                    cur = cur.Next
                }
            }

            break
        }

        cur = cur.Next
    }

    cur = head
    next := head.Next
    min = min.Next

    if cur.Next == nil && min != nil {
        min.Next = cur
        return min
    } else {

        for cur.Next != nil && min != nil {

            if cur.Val < min.Val && next.Val > min.Val {
                cur.Next = min
                minCur.Next = next
                break
            }
            cur = cur.Next
            next = next.Next
        }
    }
    return head
}

/*
   54. 螺旋矩阵

   给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

   示例 1:

   输入:
   [
    [ 1, 2, 3 ],
    [ 4, 5, 6 ],
    [ 7, 8, 9 ]
   ]
   输出: [1,2,3,6,9,8,7,4,5]
   示例 2:

   输入:
   [
     [1, 2, 3, 4],
     [5, 6, 7, 8],
     [9,10,11,12]
   ]
   输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

func spiralOrder(matrix [][]int) []int {

    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return nil
    }

    var up, down, left, right bool
    var row, column, maxRow, maxColumn, minColumn, minRow int
    maxRow = len(matrix) - 1
    maxColumn = len(matrix[0]) - 1
    var result []int

    right = true
    var cnt int
    for cnt < len(matrix)*len(matrix[0]) {
        switch {
        case left:
            result = append(result, matrix[row][column])
            if column <= minColumn {
                left = false
                up = true
                minColumn++
                minRow++
                row--
            } else {
                column--
            }

        case right:
            result = append(result, matrix[row][column])
            if column >= maxColumn {
                right = false
                down = true
                row++
            } else {
                column++
            }

        case down:
            result = append(result, matrix[row][column])
            if row >= maxRow {
                down = false
                left = true
                maxColumn--
                column--
            } else {
                row++
            }
        case up:
            result = append(result, matrix[row][column])
            if row <= minRow {
                up = false
                right = true
                column++
            } else {
                row--
            }
        }
        cnt++
    }

    return result
}

/*
   109. 有序链表转换二叉搜索树

   给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

   本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

   示例:

   给定的有序链表： [-10, -3, 0, 5, 9],

   一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

         0
        / \
      -3   9
      /   /
    -10  5
*/

func sortedListToBST(head *ListNode) *TreeNode {

    if head == nil {
        return nil
    }

    var arr []int

    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }

    left := arr[:len(arr)/2]
    var leftTree, rightTree *TreeNode
    for i := len(left) - 1; i >= 0; i-- {
        if leftTree == nil {
            leftTree = &TreeNode{Val: left[i]}
        } else {
            InsertTree(leftTree, left[i])
        }
    }

    if len(arr)/2+1 > 0 {
        right := arr[len(arr)/2+1:]

        for i := len(right) - 1; i >= 0; i-- {
            if rightTree == nil {
                rightTree = &TreeNode{Val: right[i]}
            } else {
                InsertTree(rightTree, right[i])
            }
        }
    }

    return &TreeNode{Val: arr[len(arr)/2], Left: leftTree, Right: rightTree,}
}

func InsertTree(tree *TreeNode, val int) {

    for tree != nil {
        if tree.Val < val {

            if tree.Right != nil {
                tree = tree.Right
            } else {
                tree.Right = &TreeNode{Val: val}
                return
            }
        } else {
            if tree.Left != nil {
                tree = tree.Left
            } else {
                tree.Left = &TreeNode{Val: val}
                return
            }

        }
    }

}
