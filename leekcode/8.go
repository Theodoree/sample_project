package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	max := math.MaxInt32
	min := math.MinInt32

	strs := strings.Split(strings.TrimSpace(str), ``)

	var tmp string
	var i int
	if strs[0] == `+` {
		strs = strs[1:]
	}
	for i < len(strs)  {
		if (strs[i] >= `0` && strs[i] <= `9`) || (i == 0 && strs[i] == "-") {
			tmp += strs[i]
		} else {
			break
		}
		i++
	}

	v, _ := strconv.ParseInt(tmp, 10, 64)

	if v > 0 {
		if v > int64(max) {
			return max
		}
		return int(v)
	} else {
		if v < int64(min) {
			return min
		}
		return int(v)

	}

}

func main() {

	fmt.Println(myAtoi(`+1`))

}
