package _00_899



/*
852. 山脉数组的峰顶索引

我们把符合下列属性的数组 A 称作山脉：

A.length >= 3
存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1] 的 i 的值。
*/

func peakIndexInMountainArray(A []int) int {

    var left, right int
    right = len(A) - 1
    /* 思路就是:二分查找先找到最高的位置  */
    for left <= right {

        mid := (right+left)/2
        if A[mid] > A[mid+1] && A[mid] > A[mid-1] {
            return mid
        }else if A[mid] > A[mid+1] && A[mid] < A[mid-1]{
            right = mid - 1
        }else{
            left = mid + 1
        }

    }

    return 0

}

