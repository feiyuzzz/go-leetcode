package hot100

import (
	common "github.com/feiyuzzz/go-leetcode/common"
)

func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	preHead := &common.ListNode{}
	cur := preHead

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
		} else if list1 != nil {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
	}
	return preHead.Next
}
