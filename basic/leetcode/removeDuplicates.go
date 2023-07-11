package leetcode

// 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[c] {
			c++
			nums[c] = nums[i]
		}
	}
	return c + 1
}
func Gorun() {
	removeDuplicates([]int{1, 1, 2, 2, 3, 3, 4})
}
