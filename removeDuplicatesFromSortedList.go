package demo

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, working := head, head.Next
	for working != nil {
		if pre.Val == working.Val {
			pre.Next = working.Next
		} else {
			pre = working
		}
		working = working.Next
	}
	return head
}
