/*
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25


提示：

-100.0 < x < 100.0
-231 <= n <= 231-1
n 是一个整数
-104 <= xn <= 104

题目分析：采用分治的思想，分治即将原问题分解为大于一个子问题，这里 x 的 n 次方，可以分解为 x 的 （n/2 * 2） 次方，如果 n % 2 余1则表示最后需要多乘以一个 x，否则就是 2 的整数倍，即 x 的 n 次方就是计算 n/2 次 x 的平方

x^2 * x^2 * x^2 ... = x^n
==> x^(2+2+2+2+...+2) = x^n
==> 2+2+2+2+...+2 = n
==> 一共有 n/2 个 2

则一共有 n/2 个 x^2 相乘
*/
func myPow(x float64, n int) float64 {
	// 边界条件，如果 n >= 0 的话，直接进入一般递归
	if n >= 0 {
		return quickMul(x, n)
	}
	// 否则即 n < 0 ，指数小于0，数学里面表示 n 次幂分之一，提前加 - 取反，最后结果再用1 除
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, N int) float64 {
	ans := 1.0
	x_contribute := x
	// 贡献的初始值为 x
	for N > 0 {
		// 在对 N 进行二进制拆分的同时计算答案
		if N%2 == 1 {
			//如果 N 对 2 取模余 1，那表示还需要再乘以一个 x
			ans *= x_contribute
		}
		// 否则就是 2 的倍数，则 x * x = x 的二次方
		x_contribute *= x_contribute
		// N = N /2 是一个整数，表示循环 N/2 次，及计算 N/2次 x 的二次方
		N /= 2
	}
	return ans
}