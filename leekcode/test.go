package main

/*
347. 前 K 个高频元素

给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
说明：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
*/

type heap struct {
    head *node
    cap  int
    cur  int
}

type node struct {
    left  *node
    right *node
    key   int
    val   int
}

func (h *heap) insert(n *node) {
    cur := h.head
    if h.cur == h.cap {
        if n.val < cur.val {
        
        }
        
        return
    }
}

func (h *heap) topN(n int) []int {

}

// 小顶堆
func topKFrequent(nums []int, k int) []int {
    
    var m = make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    
    h := heap{
        head: nil,
        cap:  k,
    }
    
    for k, v := range m {
        h.insert(&node{key: k, val: v})
    }
    return  h.topN(k)
}

func main() {
    c := []int{1, 2, 3, 4, 5}
}
