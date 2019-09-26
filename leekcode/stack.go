package main

import (
    . "github.com/Theodoree/sample_project/leekcode/utils"
)

func main() {

    t, _ := CreateTree([]int{5, 2, 6, 1, 3})

    PrintTree(t, "DLR", nil)

}
