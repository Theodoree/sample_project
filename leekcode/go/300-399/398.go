package _00_399



/*
398. 随机数索引

给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。

注意：
数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。

示例:

int[] nums = new int[] {1,2,3,3,3};
Solution solution = new Solution(nums);

// pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
solution.pick(3);

// pick(1) 应该返回 0。因为只有nums[0]等于1。
solution.pick(1);
*/
type Solution struct {
    nums  []int
    point int
}

func Constructor(nums []int) Solution {
    return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {

    for {
        this.point++
        if target == this.nums[this.point%len(this.nums)] {
            return this.point % len(this.nums)
        }
        if this.point >= len(this.nums) {
            this.point = 0
        }
    }

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

