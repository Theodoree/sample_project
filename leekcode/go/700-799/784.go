package _00_799

import (
	"fmt"
	"strings"
)

//784. 字母大小写全排列

func letterCasePermutation(S string) []string {

	var ret []string
	LetterCasePermutation(&ret, S, "")
	return ret

}

func LetterCasePermutation(ret *[]string, ans string, temp string) {
	if len(ans) == 0 {
		*ret = append(*ret, temp)
	}
	for i := 0; i < len(ans); i++ {
		news:=ans[:i]+ans[i+1:]
		LetterCasePermutation(ret,news,temp+strings.ToUpper(string(ans[i])))
	}

}

func main() {

	fmt.Println(letterCasePermutation(`a1b2`))

}
