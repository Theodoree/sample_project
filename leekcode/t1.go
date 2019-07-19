package main

import (
    "fmt"
)

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

func checkPossibility(nums []int) bool {

    var first int
    next := 1
    index := - 1
    for next < len(nums)-1 {

        if nums[first] > next {
            if index == -1 {
                index = first
            }
            if first != index {
                return false
            }
        }

        first++
        next++

    }

    return true

}

/*
840. 矩阵中的幻方

3 x 3 的幻方是一个填充有从 1 到 9 的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。

给定一个由整数组成的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。
输入: [[4,3,8,4],
      [9,5,1,9],
      [2,7,6,2]]
输出: 1
解释:
下面的子矩阵是一个 3 x 3 的幻方：
    438
    951
    276


*/
func numMagicSquaresInside(grid [][]int) int {

    row := 0
    index := 0
    cnt := 0
    for row < len(grid)-2 {
        sum := []int{}
        fmt.Println(row, index)
        if index < len(grid[0])-2 {
            sum = append(sum, grid[row][index]+grid[row+1][index]+grid[row+2][index])
            sum = append(sum, grid[row][index+1]+grid[row+1][index+1]+grid[row+2][index+1])
            sum = append(sum, grid[row][index+2]+grid[row+1][index+2]+grid[row+2][index+2])
        } else {
            index = 0
            row++
            continue
        }
        if sum[0] == sum[1] && sum[0] == sum[2] {
            cnt++
            index += 2
            row += 2
        }

        index++
    }

    return cnt
}

func findUnsortedSubarray(nums []int) int {

    var left, right int
    right = len(nums) - 1

    for left+1 < len(nums) && nums[left] <= nums[left+1] {
        left++
    }

    for right-1 >= 0 && nums[right-1] < nums[right] {
        right--
    }

    if left == len(nums)-1 {
        return 0
    }

    return right - left + 1
}

/*
849. 到最近的人的最大距离

在一排座位（ seats）中，1 代表有人坐在座位上，0 代表座位上是空的。

至少有一个空座位，且至少有一人坐在座位上。

亚历克斯希望坐在一个能够使他与离他最近的人之间的距离达到最大化的座位上。

返回他到离他最近的人的最大距离。

示例 1：

输入：[1,0,0,0,1,0,1]
输出：2
解释：
如果亚历克斯坐在第二个空位（seats[2]）上，他到离他最近的人的距离为 2 。
如果亚历克斯坐在其它任何一个空位上，他到离他最近的人的距离为 1 。
因此，他到离他最近的人的最大距离是 2 。
示例 2：

输入：[1,0,0,0]
输出：3

解释：
如果亚历克斯坐在最后一个座位上，他离最近的人有 3 个座位远。
这是可能的最大距离，所以答案是 3 。
*/
func maxDistToClosest(seats []int) int {

    return 0
}

//func numPairsDivisibleBy60(time []int) int {
//    var result int
//
//    for i := 0; i < len(time); i++ {
//        if time[i]%60 == 0 {
//            result++
//        } else {
//            for j := i + 1; j < len(time); j++ {
//                if j == i {
//                    continue
//                }
//                if (time[i]+time[j])%60 == 0 {
//                    fmt.Println(time[i],time[j])
//                    result++
//                }
//            }
//        }
//
//    }
//
//    return result
//}

func maximumProduct(nums []int) int {

    var one, two, third, sum int
    var min, secondmin int
    sum = 1
    one = -1e4
    two = -1e4
    third = -1e4

    for i := 0; i < len(nums); i++ {
        if nums[i] > one {
            one = nums[i]
        }

        if two < one {
            two, one = one, two
        }

        if third < two {
            third, two = two, third
        }

        if nums[i] < secondmin {
            secondmin = nums[i]
        }

        if secondmin < min {
            secondmin, min = min, secondmin
        }

    }


    if secondmin < 0 && min < 0 {
        secondmin *= -1
        min *= -1
    }
    if one != -1e4 && two != -1e4 {

        if min > two && secondmin > one {
            two = min
            one = secondmin
        }
        sum *= one
        sum *= two
    }
    /*-983 978 960 930 -925 923 -905*/
    fmt.Println(one,two,third)
    fmt.Println(one*two*third)
    if third != -1e4 {
        sum *= third
    }
    return sum
}
func main() {


    /*
    v := []int{903,606,48,-474,313,-672,872,-833,899,-629,558,-368,231,621,716,-41,-418,204,-1,883,431,810,452,-801,19,978,542,930,85,544,-784,-346,923,224,-533,-473,499,-439,-925,171,-53,247,373,898,700,406,-328,-468,95,-110,-102,-719,-983,776,412,-317,606,33,-584,-261,761,-351,-300,825,224,382,-410,335,187,880,-762,503,289,-690,117,-742,713,280,-781,447,227,-579,-845,-526,-403,-714,-154,960,-677,805,230,591,442,-458,-905,832,-285,511,536,-86}


    fmt.Println(maximumProduct(v))
    /*
       v := []int{1, 2, 3, 4}
       fmt.Println(findUnsortedSubarray(v))
       f := []int{1, 2, 3, 3, 3}
       fmt.Println(findUnsortedSubarray(f))
       f = []int{2, 6, 4, 8, 10, 9, 15}
       fmt.Println(findUnsortedSubarray(f))
       f = []int{1,3,2,2,2}
       fmt.Println(findUnsortedSubarray(f))

          v := [][]int{
              {4, 3, 8, 4},
              {9, 5, 1, 9},
              {2, 7, 6, 2},}
          fmt.Println(numMagicSquaresInside(v))

             v := []int{4, 2, 3}
             f := []int{3, 4, 2, 3}
             fmt.Println(checkPossibility(v))
             fmt.Println(checkPossibility(f))
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
