package main

import "fmt"

func commonChars(A []string) []string {

	 hashMap:=make(map[int32]int)

	for _,val:=range A{
		for _,v:=range val{
			hashMap[v]++
		}
	}


	for key,value:=range A{

		if
	}


}


func main(){

	fmt.Println(commonChars([]string{"bella","label","roller"}))

}