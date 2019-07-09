package __99

import "fmt"

//77.组合 回溯

func combine(n int, k int) [][]int {
    resultMap := make(map[string][]int)
    Combine(n, k, []int, resultMap)
}

func Combine(n int, k int, current []int, result map[string][]int) {
    if len(current) == k {

    }

}
