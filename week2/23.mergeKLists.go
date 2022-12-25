/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]


提示：
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeKList(lists, 0, len(lists1)-1)
}

func mergeKList(lists []*ListNode, left int, right int) *ListNode {
	if right == left {
		return lists[left]
	}

	mid := (left + right) / 2
	list1 := mergeKList(lists, left, mid)
	list2 := mergeKList(lists, mid+1, right)
	return mergeTwoList(list1, list2)
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	// 定义头结点
	head := new(ListNode)
	// 将头结点赋值给 node
	node := head
	// 如果 list1 不等于空并且 list2 不等于空
	for list1 != nil && list2 != nil {
		// 如果 list1 节点的值小于 list2 节点的值
		if list1.Val < list2.Val {
			// 头部节点的 Next 指向 list1
			node.Next = list1
			// list1 的 Next 赋值给 list1，进入下一轮 for 循环，即判断 list1.Next.Val 和 list2.Val 的大小
			list1 = list1.Next
			// 如果 list1 节点的值大于 list2 节点的值
		} else {
			// 则将头部节点的 Next 指向 list2
			node.Next = list2
			// 将 list2 的 Next 赋值给 list2，进入下一轮 for 循环，即判断 list2.Next.Val 和 list1.Val 的小于
			list2 = list2.Next
		}
		// 每取出一个最小，node 指针前移一位
		node = node.Next
	}
	if list1 != nil {
		node.Next = list1
	}
	if list2 != nil {
		node.Next = list2
	}
	return head.Next
}

