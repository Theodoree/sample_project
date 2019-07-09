package _00_399

import (
    "fmt"
    "strconv"
)

type Stack1 struct {
    Next *Stack1
    Val  string
}

func (s *Stack1) Push(val string) {
    stack1 := &Stack1{Val: val}
    stack1.Next = s
    s = stack1
}

func (s *Stack1) Pop() string {
    str1 := s.Val
    s = s.Next
    return str1
}

//字符串解码
func decodeString(s string) string {

    KuohaoStack := &Stack1{}
    StrStack := &Stack1{}
    var nums []int
    var str string
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            num, _ := strconv.Atoi(string(s[i]))
            nums = append(nums, num)
        } else if string(s[i]) == "[" {
            StrStack.Push(str)
            str = ""
            KuohaoStack.Push("[")
        } else if string(s[i]) == "]" {
            StrStack.Push(str)
            str = ""
            KuohaoStack.Push("]")
        }
        str += string(s[i])

    }
    if len(str) > 0 {
        StrStack.Push(str)
    }

    var s string
    var cnt int
    for {
        if cnt == 0 {
            cnt = nums[0]
            nums = nums[1:]
        }


    }

}

func main() {

    fmt.Println(decodeString(`3[a]2[bc]`))

}
