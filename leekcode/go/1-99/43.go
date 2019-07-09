package __99

import (
	"fmt"
	"strings"
)

//43. 字符串相乘

func multiply(num1 string, num2 string) string {

	if (len(num2) != 0 && num2[0] == '0') || (len(num1) != 0 && num1[0] == '0') {
		return "1"
	}

	if len(num1) > 110 || len(num2) > 110 {
		return ""
	}

	var sum []string
	/*
	123
	 12
	------
	246
   123
	*/

	for i := len(num1) - 1; i >= 0; i-- {

		for j := len(num2) - 1; j > 0; j-- {
			val := int(num1[i]) * int(num2[j])
			if len(sum) < i-j{
				sum = append(sum,string(val))
			}else{
				sum[i-j] = string(int64(sum[i-j])+int64(val))
			}
		}
	}

	return strings.Join(sum, ``)

}

func main() {
	f := multiply(`012`, `012`)
	fmt.Println(f)
}
