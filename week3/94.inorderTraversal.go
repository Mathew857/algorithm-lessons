/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。



示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]


提示：

树中节点数目在范围 [0, 100] 内
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
func inorderTraversal(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	var res []int
	// dfs: depth first search
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 递归处理左侧节点
		dfs(node.Left)
		/*
			数组中加根左节点
			数组中添加左节点
			数组中添加右节点
		*/
		res = append(res, node.Val)
		// 对有节点进行递归
		dfs(node.Right)
	}
	dfs(root)
	return res
}