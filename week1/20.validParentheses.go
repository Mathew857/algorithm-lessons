/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
*/
func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}
	stack := make([]rune, len(s))
	i := 0
	for _, value := range s {
		switch value {
		case '(', '[', '{':
			stack[i] = value
			i++
		case ')':
			if i-1 >= 0 && stack[i-1] == '(' {
				i--
			} else {
				return false
			}
		case ']':
			if i-1 >= 0 && stack[i-1] == '[' {
				i--
			} else {
				return false
			}
		case '}':
			if i-1 >= 0 && stack[i-1] == '{' {
				i--
			} else {
				return false
			}
		}
	}
	if i == 0 {
		return true
	} else {
		return false
	}
}