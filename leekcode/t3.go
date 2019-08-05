package main

import "fmt"

/*
420. 强密码检验器

一个强密码应满足以下所有条件：

由至少6个，至多20个字符组成。
至少包含一个小写字母，一个大写字母，和一个数字。
同一字符不能连续出现三次 (比如 "...aaa..." 是不允许的, 但是 "...aa...a..." 是可以的)。
编写函数 strongPasswordChecker(s)，s 代表输入字符串，如果 s 已经符合强密码条件，则返回0；否则返回要将 s 修改为满足强密码条件的字符串所需要进行修改的最小步数。

插入、删除、替换任一字符都算作一次修改。
*/

func strongPasswordChecker(s string) int {
    /*
       1.由至少6个，至多20个字符组成。
       2.至少包含一个小写字母，一个大写字母，和一个数字。
       3.同一字符不能连续出现三次
    */
    var Num, Lowercase, M bool
    var result int
    var cnt int
    b := []byte(s)
    for i := 1; i < len(b); i++ {
        if b[i-1] == b[i] {
            cnt++
        } else {
            cnt = 0
        }

        switch {
        case !Num && b[i] >= '0' && b[i] <= '9':
            Num = true
        case !Lowercase && b[i] >= 'a' && b[i] <= 'z':
            Lowercase = true
        case !M && b[i] >= 'A' && b[i] <= 'Z':
            M = true
        }

        if cnt == 2 {
            b[i] = b[i] - 1
            result++
            cnt = 0
        }
    }

    for len(b) > 20 {
        result++
        b = b[1:]
    }

    for i := 0; i < result; i++ {
        switch {
        case !Num:
            Num = true
        case !Lowercase:
            Lowercase = true
        case !M:
            M = true

        }

    }

    if !Num {
        if len(s) < 6 {
            s += "1"
        }
        result++
    }

    if !Lowercase {
        if len(s) < 6 {
            s += "1"
        }
        result++
    }

    if !M {
        if len(s) < 6 {
            s += "1"
        }

        result++

    }

    for len(s) < 6 {
        result++
        s += " "
    }



    return result
}

func main() {
    fmt.Println(strongPasswordChecker("aaaaaaaaaaaaaaaaaaaaa"))
    /*
       12009.0,
       6987.0,
       4344.66667,
       1129.5,
       -2385.6,
       -3213.4,
       1502.6,
       -2515.0,
       2874.2,
       1470.6
    */

    /*
       DR!DDDLD!D!
       DR!DDDLD!!

    */
}
