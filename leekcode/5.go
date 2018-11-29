package main

import (
	"strings"
	"sort"
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) > 1000 {
		return ""
	}
	sd := strings.Split(s, "")
	df := make(map[int][]int)
	length_list := []int{}
	for i := 1; i < len(sd); i++ {
		var index, end, length int
		length = 1
		for i-length >= 0 && i+length < len(sd) {
			if sd[i-length] == sd[i+length] {
				index = i - length
				end = i + length
				length += 1
			} else {
				df[length] = []int{index, end}
				length_list = append(length_list, length)
				break
			}
		}
	}
	sort.Ints(length_list)
	value := df[length_list[len(length_list)-1]]
	str := strings.Join(sd[value[0]:value[1]+1], "")
	return str
}

func main() {
	fmt.Println(longestPalindrome("babaddsfdsf"))

}
