package hot100

import "github.com/feiyuzzz/go-leetcode/common"

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	head := &common.ListNode{}
	cur := head
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		if head == nil {
			head = &common.ListNode{Val: sum % 10}
			head = head.Next
		} else {
			cur = &common.ListNode{Val: sum % 10}
			cur = cur.Next
		}
	}
	if carry > 0 {
		cur.Next = &common.ListNode{Val: carry}
	}
	return head.Next
}
