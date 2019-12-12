package main

import (
    "fmt"
)

func main() {
    
    v := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(v)-1; i++ {
        fmt.Println(v[i+1])
    }
    
}
