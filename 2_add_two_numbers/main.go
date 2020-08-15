package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{}
	var p = l1
	var q = l2
	var current = result
	var carry int
	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		var sum = carry + x + y
		carry = sum / 10
		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}
	return result.Next

}
