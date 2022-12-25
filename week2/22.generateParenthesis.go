/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]


提示：

1 <= n <= 8
*/

package main

func generateParenthesis(n int) []string {
	vret := make([]string, 0)
	getAbc(n, n, &vret, "", 0)
	return vret
}

// left:  左括号("(") 当前剩余个数
// right: 右括号(")") 当前剩余个数
func getAbc(left, right int, vret *[]string, str string, cnt int) {
	//	fmt.Printf("vret is %v\n", vret)
	// 左括号剩余的个数应该 小与 右括号剩余的个数
	if left > right {
		return
	}
	if right == 0 && left == 0 {
		*vret = append(*vret, str)
		// vret = append([]string{}, "((()))")
	}
	/*
		注意这里每次 getAbc，left > 0 和 right > 0 都会执行到
	*/
	if left > 0 {
		//fmt.Printf("第 %v 次调用，left  > 0, left is %v, right is %v, call getAbc(%v, %v, %v, %v)\n", cnt+1, left, right, left-1, right, *vret, str+"(")
		getAbc(left-1, right, vret, str+"(", cnt+1)
		/*
			getAbc(2, 3, []string{}, "(")
				getAbc(1, 3, []string{}, "((")
					getAbc(0, 3, []string{}, "(((")
		*/

	}
	if right > 0 {
		//fmt.Printf("第 %v 次调用，right > 0, left is %v, right is %v, call getAbc(%v, %v, %v, %v)\n", cnt+1, left, right, left, right-1, *vret, str+")")
		getAbc(left, right-1, vret, str+")", cnt+1)
		/*
			getAbc(0, 2, []string, "((()")
				getAbc(0, 1, []string, "((())")
					getAbc(0, 0, []string, "((()))")
		*/
	}
}

func main() {
	generateParenthesis(3)
}
