package algo

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	length, tailPtr := 1, head
	for {
		if tailPtr.Next == nil {
			break
		}
		tailPtr = tailPtr.Next
		length++
	}

	headPtr := &ListNode{Next: head}
	for i := 0; i < length/2; i++ {
		node := headPtr.Next
		headPtr.Next = node.Next
		node.Next = tailPtr.Next
		tailPtr.Next = node
	}

	if length%2 == 1 {
		headPtr = headPtr.Next
	}

	for i := 0; i < length/2; i++ {
		if headPtr.Next.Val != tailPtr.Next.Val {
			return false
		}
		headPtr, tailPtr = headPtr.Next, tailPtr.Next
	}
	return true
}
