package algo

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	for i := 0; i < n; i++ {
		q = q.Next
	}

	if q == nil {
		return head.Next
	}

	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return head
}

func removeNthFromEndAnswer(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
