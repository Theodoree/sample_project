package main

import (
    "fmt"
    "strings"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//1.空字符串是有效的字符串
//2.只考虑字母和数字字符,忽略字母的大小写
func isPalindrome(s string) bool {

    if s == "" {
        return true
    }
    //A man, a plan, a canal: Panama
    s = strings.ToLower(s)
    first := 0
    last := len(s) - 1
    for first < last {

        for (s[first] < 'a' || s[first] > 'z') && (s[first] < '0' || s[first] > '9') {
            first++
            if first >= len(s) {
                return false
            }
        }
        for (s[last] < 'a' || s[last] > 'z') && (s[last] < '0' || s[last] > '9') {
            last--
            if last < 0 {
                return false
            }
        }

        if first > last {
            return true
        }

        if s[first] != s[last] {
            return false
        }
        first++
        last--
    }

    return true
}

//func isValidBST(root *TreeNode) bool {
//
//    if root == nil {
//        return true
//    }
//
//    if root.Left != nil {
//        if root.Val < root.Left.Val || root.Val == root.Left.Val {
//            return false
//        }
//
//    }
//    if root.Right != nil {
//        if root.Val > root.Right.Val || root.Val == root.Right.Val {
//            return false
//        }
//    }
//
//    return isValidBST(root.Left) == true && isValidBST(root.Right) == true
//
//}

func rob(nums []int) int {

    var sum, sum1 int

    for i := 0; i < len(nums); i++ {
        if (i+1)%2 == 0 {
            sum+=nums[i]
        }else{
            sum1+=nums[i]

        }
    }


    if sum > sum1{
        return  sum
    }

    return  sum1

}
func main() {
    fmt.Println(rob([]int{2,1,1,2}))

}
