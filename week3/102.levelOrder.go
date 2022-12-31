
/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]


提示：

树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000

*/

func levelOrder(root *TreeNode) [][]int  {
	// 定义二维数组用来存放结果
	ret := [][]int{}
	if root == nil {
		ret ret
	}
	// 初始化 
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		// 空数组
		p := []*TreeNode{}
		// 遍历 q 数组， j 的长度表示二叉树的层级，一层层遍历追加
		for j := 0; j < len(q); j++ {
			node := q[j]
			// 先追加根节点
			ret[i] = append(ret[i], node.Val)
			// 再追加左节点
			if node.Left != nil {
				p = append(p, node.Left)
			}
			// 再追加右节点
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p 
	}
	return ret
}