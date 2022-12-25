/*
网站域名 "discuss.leetcode.com" 由多个子域名组成。顶级域名为 "com" ，二级域名为 "leetcode.com" ，最低一级为 "discuss.leetcode.com" 。当访问域名 "discuss.leetcode.com" 时，同时也会隐式访问其父域名 "leetcode.com" 以及 "com" 。

计数配对域名 是遵循 "rep d1.d2.d3" 或 "rep d1.d2" 格式的一个域名表示，其中 rep 表示访问域名的次数，d1.d2.d3 为域名本身。

例如，"9001 discuss.leetcode.com" 就是一个 计数配对域名 ，表示 discuss.leetcode.com 被访问了 9001 次。
给你一个 计数配对域名 组成的数组 cpdomains ，解析得到输入中每个子域名对应的 计数配对域名 ，并以数组形式返回。可以按 任意顺序 返回答案。



示例 1：

输入：cpdomains = ["9001 discuss.leetcode.com"]
输出：["9001 leetcode.com","9001 discuss.leetcode.com","9001 com"]
解释：例子中仅包含一个网站域名："discuss.leetcode.com"。
按照前文描述，子域名 "leetcode.com" 和 "com" 都会被访问，所以它们都被访问了 9001 次。
示例 2：

输入：cpdomains = ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
输出：["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
解释：按照前文描述，会访问 "google.mail.com" 900 次，"yahoo.com" 50 次，"intel.mail.com" 1 次，"wiki.org" 5 次。
而对于父域名，会访问 "mail.com" 900 + 1 = 901 次，"com" 900 + 50 + 1 = 951 次，和 "org" 5 次。


提示：

1 <= cpdomain.length <= 100
1 <= cpdomain[i].length <= 100
cpdomain[i] 会遵循 "repi d1i.d2i.d3i" 或 "repi d1i.d2i" 格式
repi 是范围 [1, 104] 内的一个整数
d1i、d2i 和 d3i 由小写英文字母组成
*/

package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	cnt := map[string]int{}
	// 循环 cpdomains
	for _, s := range cpdomains {
		// 返回空格的下标
		i := strings.IndexByte(s, ' ')
		// 将 9001 转换为 int
		c, _ := strconv.Atoi(s[:i])
		// 取出域名
		s = s[i+1:]
		// key 为 数组元素，value 为
		// cnt[s] = cnt[s] // 初始默认为 0  + c， 0 + c 为 c，即表示把 9001 赋值给 cnt["discuss.leetcode.com"]
		cnt[s] += c
		for {
			// 取出域名第一个 . 的索引
			/*
				第一次循环 s 为 discuss.leetcode.com， i 为 7
				第二次循环 s 为 leetcode.com，i 为 8
				第三次循环 s 为 com ， 因为没有 . 所以返回 -1，i 为 -1
			*/
			i := strings.IndexByte(s, '.')
			if i < 0 {
				break
			}
			// 原 s 为 discuss.leetcode.com ，s[i+1:] 为 leetcode.com
			/*
				第一次循环得到 s = ["leetcode.com"]
				第二次循环得到 s = ["com"]
			*/
			s = s[i+1:]
			// cnt["leetcode.com"] = 9001
			/*
				第一次循环得到 cnt["leetcode.com"] = 9001
				第二次循环得到 cnt["com"] = 9001
				此时 cnt 中公有三个 key ，分别为 discuss.leetcode.com leetcode.com com , value 均为 9001
			*/
			cnt[s] += c
		}
		// 退出循环
	}
	// 构造长度为 cnt 长度的数组，初始值为0
	ans := make([]string, 0, len(cnt))
	// 遍历 cnt map
	for s, c := range cnt {
		// 数组追加 9001 转字符，加上 map key 为子域名，即:
		/*
			ans = ["9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"]
		*/
		ans = append(ans, strconv.Itoa(c)+" "+s)
	}
	return ans
}

func main() {
	cpdomains := []string{"9001 discuss.leetcode.com"}
	subdomainVisits(cpdomains)
}
