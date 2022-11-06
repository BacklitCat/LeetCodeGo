package algo

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := ListNode{0, nil}
	r, p, q := &root, list1, list2
	for p != nil && q != nil {
		if p.Val < q.Val {
			r.Next = p
			p = p.Next
		} else {
			r.Next = q
			q = q.Next
		}
		r = r.Next
	}
	for p != nil {
		r.Next = p
		p = p.Next
		r = r.Next
	}
	for q != nil {
		r.Next = q
		q = q.Next
		r = r.Next
	}

	p = If3(p != nil, p, q).(*ListNode)
	for p != nil {
		r.Next = p
		p = p.Next
		r = r.Next
	}

	return root.Next
}
