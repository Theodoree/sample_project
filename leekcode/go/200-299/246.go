package _00_299


/*
246. 中心对称数

中心对称数是指一个数字在旋转了 180 度之后看起来依旧相同的数字（或者上下颠倒地看）。

请写一个函数来判断该数字是否是中心对称数，其输入将会以一个字符串的形式来表达数字。

示例 1:

输入:  "69"
输出: true
示例 2:

输入:  "88"
输出: true
示例 3:

输入:  "962"
输出: false
*/

func isStrobogrammatic(num string) bool {
    m := map[byte]byte{
        '0': '0',
        '1': '1',
        '8': '8',
        '6': '9',
        '9': '6',
    }
    var result []byte
    for i := len(num) - 1; i >= 0; i-- {
        result = append(result, m[num[i]])
    }
    return string(result) == num
}
