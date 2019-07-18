package main

func hasGroupsSizeX(deck []int) bool {

    if len(deck) == 1 {
        return false
    }
    filter := make(map[int]int)
    for _, v := range deck {
        filter[v]++
    }

    var last int
    for _, v := range filter {
        if v == 1 {
            return false
        }
        if last == 0 {
            last = v
            for last%2 == 0 && last/2 > 1 {
                last /= 2
            }

        }

        switch {
        case last > v:
            if last%v != 0 {
                return false
            }
        case last < v:
            if v%last != 0 {
                return false
            }
        }

    }
    return true
}



func main() {

    /*
       fmt.Println(hasGroupsSizeX([]int{
       //    0,0,0,0,
       //    0,0,0,0,
       //    0,0,0,0,
       //    0,0,
       //    1,1,1,1,
       //    1,1,
       //    2,2,2,2,
       //    2,2,
       //    3,3,3,3,
       //    3,3,3,3
       //    ,3,3,3,3,
       //    4,4,
       //    5,5,5,5,
       //    6,6,
       //    7,7,
       //    8,8}))
    */

}
