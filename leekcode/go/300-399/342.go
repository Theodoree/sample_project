package _00_399


/*

342. 4的幂

给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。

示例 1:

输入: 16
输出: true
示例 2:

输入: 5
输出: false

1024  10000000000
256   00100000000
64    00001000000
16    00000001000
4     00000000010
2048 1024 512 256 128 64 32 16 8 4 2 1
     1024     256     64    16   4
*/

func isPowerOfFour(num int) bool {

    n := 1

    for n <= num {
        if n == num {
            return true
        }
        n <<=2
    }

    return false
}
