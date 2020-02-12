package _go



/*
面试题66. 构建乘积数组

给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。



示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]


提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000
在真实的面试中遇到过这道题？
*/

func constructArr(a []int) []int {
    n := len(a)
    if n == 0 {
        return []int{}
    }
    b := make([]int, n)
    left := 1
    for i := 0; i < n; i++ {
        b[i] = left
        left *= a[i]
    }
    right := a[n-1]
    for i := n - 2; i >= 0; i-- {
        b[i] *= right
        right *= a[i]
    }
    return b
}