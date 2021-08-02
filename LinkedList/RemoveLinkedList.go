/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	for head.Val == val {
		head = head.Next
		if head == nil {
			return head
		}
	}

	curr := head
	var prev *ListNode

	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			curr = curr.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return head
}