package __99

//39.组合总和
func combinationSum(candidates []int, target int) [][]int {

	var res [][]int

	CombinationSum(&res, candidates, []int{}, target)

	return res
}

func CombinationSum(res *[][]int, candidates []int, temp []int, target int) {

	if sum(temp) == target {
		*res = append(*res, temp)
		return
	}

}

func sum(slice []int) int {

	var sum int
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum
}
