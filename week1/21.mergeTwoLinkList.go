
/*
* Definition for singly-linked list.
* type ListNode struct {
*      Val int
*      Next *ListNode
* }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 如果 list1 等于 list2 ，优先返回 list2
	if list1 == nil {
		return list2
	}

	// 如果 list2 等于空，则返回 list1
	if list2 == nil {
		return list1
	}

	if list1 == list2 {
		return list2
	}

	// 判断 list1 和 list2 哪一个链表的头结点的值更小，然后递归决定下一个添加到结构里面的节点。
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

/*
//(1,1):代表第一次进入递归函数，并且从第一个口进入，并且记录进入前链表的状态
merge(1,1): 1->4->5->null, 1->2->3->6->null
    merge(2,2): 4->5->null, 1->2->3->6->null
    	merge(3,2): 4->5->null, 2->3->6->null
    		merge(4,2): 4->5->null, 3->6->null
    			merge(5,1): 4->5->null, 6->null
    				merge(6,1): 5->null, 6->null
    					merge(7): null, 6->null
    					return l2
    				l1.next --- 5->6->null, return l1
    			l1.next --- 4->5->6->null, return l1
    		l2.next --- 3->4->5->6->null, return l2
    	l2.next --- 2->3->4->5->6->null, return l2
    l2.next --- 1->2->3->4->5->6->null, return l2
l1.next --- 1->1->2->3->4->5->6->null, return l1
*/