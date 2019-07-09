package _00_999

import "fmt"

type StockSpanner struct {
	prices []int
}

func Constructor() StockSpanner {
	return StockSpanner{prices: []int{}}
}

func (this *StockSpanner) Next(price int) int {
	this.prices = append(this.prices, price)
	if len(this.prices) == 1 {
		return 1
	}

	cnt := 1
	for i := len(this.prices) - 2; i >= 0; i-- {
		if this.prices[i] > price {
			break
		}
		cnt++
	}

	return  cnt

}

func main(){

	S := Constructor()
	var c  int
	c=S.Next(100) //被调用并返回 1，
	fmt.Println(c)
	c=S.Next(80) //被调用并返回 1，
	fmt.Println(c)

	c=S.Next(60) //被调用并返回 1，
	fmt.Println(c)

	c=S.Next(70) //被调用并返回 2，
	fmt.Println(c)

	c=S.Next(60) //被调用并返回 1，
	fmt.Println(c)

	c=S.Next(75) //被调用并返回 4，
	fmt.Println(c)

	c=S.Next(85) //被调用并返回 6。
	fmt.Println(c)



}
