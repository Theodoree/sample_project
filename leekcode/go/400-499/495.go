package _00_499

import (
	"fmt"
	"math"
)

//495. 提莫攻击

func findPoisonedDuration(timeSeries []int, duration int) int {

	var cnt int
	cnt = duration
	for i := 0; i < len(timeSeries)-1; i++ {
		cnt+=int(math.Min(float64(timeSeries[i+1]-timeSeries[i]),float64(duration)))
	}

	if cnt > 10000000 {
		cnt = 10000000

	}

	return cnt
}

func main(){

	fmt.Println(findPoisonedDuration([]int{1,4,5,6},2))


}
