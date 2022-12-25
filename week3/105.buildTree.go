/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:
输入: preorder = [-1], inorder = [-1]
输出: [-1]

提示:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder 和 inorder 均 无重复 元素
inorder 均出现在 preorder
preorder 保证 为二叉树的前序遍历序列
inorder 保证 为二叉树的中序遍历序列
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 如果前序节点数组长度为空，直接返回 nil
	if len(preorder) == 0 {
		return nil
	}
	// 构造根节点，根节点即为前序数组 0 号元素
	root := &TreeNode{preorder[0], nil, nil}
	// 定义 i
	i := 0
	for ; i < len(inorder); i++ {
		// 取出中序数组中根节点的索引位置 i
		if inorder[i] == preorder[0] {
			break
		}
	}
	// len(inorder[:i]) 表示左子树的节点个数
	// inorder[:i] 表示左子树的元素数组，len(inorder[:])+1
	// 根节点的左子树 等于 buildTree(左子树和右子树，)
	// buildTree(左子树数组，根节点+左子树 数组)
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	// inorder[i+1:] 表示右子树数组
	// preorder[len(inorder[:i])+1:] 表示取出 根节点+左子树 数组
	// buildTree(根节点+左子树 数组，右子树数组)
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

