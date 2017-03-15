package demo

func hasCycle(ListNode *ListNode) bool {
	if ListNode == nil || ListNode.Next == nil {
		return false
	}
	slow, fast := ListNode, ListNode.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// fast will catch slow if linked list is cycle
		if slow == fast {
			return true
		}

	}

	return false
}
