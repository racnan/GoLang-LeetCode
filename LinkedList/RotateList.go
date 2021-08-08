/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	curr := head
	var lent = 0

	for curr != nil {
		lent++
		curr = curr.Next
	}

	if lent == 0 || lent == 1 {
		return head
	}

	if lent <= k {
		k = k % lent
	}

	if k == 0 {
		return head
	}

	k = lent - k
	curr = head
	var prev *ListNode

	for k != 0 {
		k--
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil
	newHead := curr

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = head

	return newHead
}