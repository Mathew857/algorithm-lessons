/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”



示例 1：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
示例 2：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
示例 3：

输入：root = [1,2], p = 1, q = 2
输出：1
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果根节点为空直接返回空
	if root == nil {
		return nil
	}
	// 如果根节点的值等 p 的值或者 根节点的值等于 q 的值，则直接返回根节点
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 进入递归判断左孩子节点
	// 如果根节点的左节点 等于 p 或者根节点的左节点 等于 q ，则直接返回 根节点的左孩子节点
	left := lowestCommonAncestor(root.Left, p, q)
	// 如果根节点的右节点 等于 p 或者根节点的右节点 等于 q ，则直接返回 根节点的右孩子节点
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}