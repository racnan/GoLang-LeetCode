/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var head ListNode
	var curr *ListNode

	curr = &head

	for true {
		if l1 == nil {
			curr.Next = l2
			break
		} else if l2 == nil {
			curr.Next = l1
			break
		}

		if l1.Val < l2.Val {
			curr.Next = l1
			curr = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			curr = l2
			l2 = l2.Next
		}
	}
	return head.Next
}