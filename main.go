package main

import (
	"fmt"
)

func main(){

	a:=GetPercentList([]int{5169,17,30,13,6,2,6,2,6,12,2,2,1,1,2,2,1})
	fmt.Println(a)
	for _,value:=range a{
		fmt.Print(value,",")
	}
}


func GetPercentList(valueList []int)(PercentList []string){
	var total float64
	for _,value:=range valueList{
		total+=float64(value)
	}

	for _,value:=range valueList{
		percent:=float64(value)/total
		fmt.Println(percent)
		PercentList = append(PercentList,fmt.Sprintf("%.2f ",percent*100))
	}


	return
}