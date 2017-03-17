package demo

func removeLinkedListElements(head *ListNode, val int) *ListNode {
	newHead := newListNode(-1)
	newHead.Next = head

	working := newHead.Next
	pre := newHead
	for working != nil {
		if working.Val == val {
			//if target locates at beging, pre.Next is newHead.Next
			//or
			//pre moves forward to handle all nodes after newHead.Next
			pre.Next = working.Next
			working = working.Next
			continue
		}
		pre = working
		working = working.Next
	}
	return newHead.Next
}
