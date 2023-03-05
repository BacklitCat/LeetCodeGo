package algo

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var slow, fast, ptr = head, head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // has cycle
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}
	return nil
}
