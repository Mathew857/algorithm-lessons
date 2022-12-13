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

	var LastNode *ListNode
	for head != nil {
		next := head.Next
		head.Next = LastNode
		LastNode = head
		head = next
	}
	return LastNode
}