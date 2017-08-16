package demo

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode
	var travel *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		ret = l2
		l2 = l2.Next
	} else {
		ret = l1
		l1 = l1.Next
	}
	travel = ret

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			travel.Next = l2
			l2 = l2.Next
		} else {
			travel.Next = l1
			l1 = l1.Next
		}
		travel = travel.Next
	}
	for l2 != nil {
		travel.Next = l2
		l2 = l2.Next
		travel = travel.Next
	}

	for l1 != nil {
		travel.Next = l1
		l1 = l1.Next
		travel = travel.Next
	}

	return ret
}
