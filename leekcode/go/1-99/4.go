package __99

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums []int, nums2 []int) float64 {
	nums = append(nums, nums2...)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	lennumber := len(nums)
	if lennumber%2 == 1 {
		return float64(nums[lennumber/2])
	} else {
		return float64(nums[lennumber/2-1]+nums[lennumber/2]) / 2.0
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
