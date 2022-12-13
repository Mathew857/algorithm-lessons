func plusOne(digits []int) []int {
	// 获取数组长度
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		// 从最后一个元素开始判断是否等于 9
		if digits[i] != 9 {
			// 不等于9的话末尾元素直接加一
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}