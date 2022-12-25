/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

 

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]

提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

题目分析：
通常情况下，默认会想到 O(n^2) 的算法，即两层 for 循环
    for i, x := range nums {
        for j := i + 1; j < len(nums); j++ {
            if x+nums[j] == target {
                return []int{i, j}
            }
        }
    }

通过 hashtable 的方式，可以将复杂度降低为 O(n)
*/
func twoSum([]int, target int) []int {
	// 通过 hashtable ，将数组元素本身存储为 hashtable 的 key，数组元素对应下标索引存储为 hashtable 的 value 
	hashtable := map[int]int{}
	// 遍历数组
	for i, v := range nums {
		// 如果 hashtable[target-v] 表示 目标值 target - v 存在数组中，则直接返回 v 索引和 p 索引
		if p, ok := hashtable[target-v]; ok {
			return []int{p, i}
		}
		// 如果不存在，则将 v 插入到 hashtable 中，等待后续循环遍历判断是否存在满足条件的元素
		hashtable[v] = i
	}
	return nil
}