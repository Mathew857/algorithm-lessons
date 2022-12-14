/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	cur := head
	for cur != nil {
		pre, cur, preNext := cur, cur.Next, pre
	}
}