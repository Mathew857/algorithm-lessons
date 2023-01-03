/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]


提示:
1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
*/

func moveZero(nums []int) {
	// 两层 for 循环，第一层取出 非0值，第二层将后续位置补0
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}
	for i := n; i < len(nums); i++ {
		nums[i] = 0
	}
}