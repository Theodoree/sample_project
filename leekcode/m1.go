package main

func DeferFunc(i int) int {
	t := i


	defer func() {
		t += 3
	}()


	return t
}



func main() {

	println(DeferFunc(10))

}
