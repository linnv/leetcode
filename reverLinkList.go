package demo

func ReverseList(head *ListNode) *ListNode {
	var pre, cur, backup *ListNode
	cur = head
	for cur != nil {
		backup = cur.Next
		cur.Next = pre
		pre = cur
		cur = backup
	}
	return pre
}
