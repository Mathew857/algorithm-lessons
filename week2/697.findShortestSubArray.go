/*
给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

示例 1：

输入：nums = [1,2,2,3,1]
输出：2
解释：
输入数组的度是 2 ，因为元素 1 和 2 的出现频数最大，均为 2 。
连续子数组里面拥有相同度的有如下所示：
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组 [2, 2] 的长度为 2 ，所以返回 2 。

示例 2：

输入：nums = [1,2,2,3,1,4,2]
输出：6
解释：
数组的度是 3 ，因为元素 2 重复出现 3 次。
所以 [2,2,3,1,4,2] 是最短子数组，因此返回 6 。


提示：

nums.length 在 1 到 50,000 范围内。
nums[i] 是一个在 0 到 49,999 范围内的整数。
*/
package main

type entry struct {
	// 出现次数、该数在原数组中第一次出现的位置、该数在数组中最后一次出现的位置
	cnt, l, r int
}

func findShortestSubArray(nums []int) (ans int) {
	// 构造 mp map，key 为 数组元素，value 为 entry，包含其出现次数，出现第一次最后一次的位置
	mp := map[int]entry{}
	// 遍历数组
	for i, v := range nums {
		// 如果 map 中存在该数
		if e, has := mp[v]; has {
			// 则 map 中该元素次数统计加一
			e.cnt++
			// 记录 最后一次出现位置为该元素在数组中的索引
			// 如果遇到重复元素，则最后一次出现的索引会被更新为较后的索引
			e.r = i
			// 更新 map[v]
			mp[v] = e
		} else {
			// 如果元素不重复，则前后索引均为第一次出现的索引
			mp[v] = entry{1, i, i}
		}
	}
	/*
		mp := {1:{2 0 4} 2:{2 1 2} 3:{1 3 3}}
	*/
	maxCnt := 0
	// map 构造完之后，遍历 map
	for _, e := range mp {
		// 如果 e 出现的次数大于 maxCnt，取出 maxCnt 和 其长度
		if e.cnt > maxCnt {
			// 更新 maxCnt 为 e.cnt，ans 为 e 的最后一次出现的索引-第一次出现的索引加一
			// 即 ans 表示第一个 e 和最后一个e 之间的长度
			// 遍历一遍之后得到 maxCnt 为 2
			maxCnt, ans = e.cnt, e.r-e.l+1
			// 如果元素出现次数和最大次数相等
		} else if e.cnt == maxCnt {
			// 则取出 索引距离较近的值，题目为返回最短连续子数数组长度
			ans = min(ans, e.r-e.l+1)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 2, 3, 1}
	findShortestSubArray(nums)
}
