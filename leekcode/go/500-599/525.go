package _00_599

import "fmt"

type stack2 struct {
    Next *stack2
    Val  int
}

func (s *stack2) push(i int) {

    stack := &stack2{Val: i}
    if s == nil {
        s = stack
        return
    }
    stack.Next = s
    s = stack

}

func (s *stack2) pop() int {
    if s == nil {
        return -1
    }

    val := s.Val
    s = s.Next

    return val
}

func (s *stack2) Clear(i int) {
    s.Next = nil
    s.Val = i
}

func findMaxLength(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    s := &stack2{Val: nums[0]}

    var maxCnt, cnt int
    cnt = 1
    for i := 1; i < len(nums); i++ {
        num := s.pop()
        fmt.Println(num,nums[i],nums[0])
        switch num {
        case 0:
            if nums[i] == 1 {
                s.push(nums[i])
                cnt++
            } else {
                if maxCnt < cnt {
                    maxCnt = cnt
                }
                cnt = 1
                s.Clear(nums[i])
            }
        case 1:
            if nums[i] == 0 {
                s.push(nums[i])
                cnt++
            } else {
                if maxCnt < cnt {
                    maxCnt = cnt
                }
                cnt = 1
                s.Clear(nums[i])
            }
        }
    }

    if cnt > maxCnt {
        maxCnt = cnt
    }

    return maxCnt
}

func main() {
    fmt.Println(findMaxLength([]int{0, 1, 1}))

}
