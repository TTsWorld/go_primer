package other

import (
	"fmt"
	"testing"
)

/*
//滴滴
    给你一个整数数组 nums ，判断是否存在三元组[nums[i], nums[j], nums[k]]
	满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
	请你返回所有和为 0 且不重复的三元组。
	注意：答案中不可以包含重复的三元组。
	输入：nums = [-1,0,1,2,-1,-4]
	输出：[[-1,-1,2],[-1,0,1]]
	解释：
	nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
	nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
	nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
	不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
	注意，输出的顺序和三元组的顺序并不重要。
*/

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		t := -nums[i]
		m := make(map[int]int, len(nums))
		for j := 0; j < len(nums); j++ {
			k, ok := m[t-nums[j]]
			if !ok {
				m[nums[j]] = j
				continue
			}
			if i == j || i == k || j == k {
				continue
			}
			ans = append(ans, []int{t, nums[j], k})
		}
	}
	return ans
}

func TestI(t *testing.T) {
	fmt.Printf("%v", threeSum([]int{-1, 0, 1, 2, -1, -4}))

}
