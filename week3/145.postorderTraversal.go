/*
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。

示例 1：
输入：root = [1,null,2,3]
输出：[3,2,1]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

提示：

树中节点的数目在范围 [0, 100] 内
-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 通过迭代模拟递归的先进后出来实现后序遍历
// 前中序遍历来说只需要改变当前节点和其左右子节点的入栈顺序即可
// 非常好的统一了三种遍历
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	// 先让根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			// 当前节点先入栈，这样出栈时就能保证比子节点后出栈
			// 并且对于已经处理过子节点的节点，会在其再次入栈的时候再压入一个 nil 作为标记
			// 下次遍历到 nil 就直接将 nil 下面的节点出栈
			stack = append(stack, top, nil)
			// 子节点入栈顺序是先右后左，保证出栈先左后右
			if top.Right != nil {
				// 右节点先入栈
				stack = append(stack, top.Right)
			}
			// 右节点入栈后左节点再入栈，这样保证左节点先出栈，右节点后出栈
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
			// 当一直 top.Right 或者 Top.Left 的时候，可能回到 nil 节点，即执行 else
		} else if len(stack) > 0 {
			// 已处理过子节点的节点直接出栈并加入输出
			// 表示把栈顶节点赋值给 top
			top = stack[len(stack)-1]
			// 在 ans 中追加栈顶节点的 value
			ans = append(ans, top.Val)
			// 栈底元素移出
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}