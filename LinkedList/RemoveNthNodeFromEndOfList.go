/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	delay := head
	var prevOfDelay *ListNode

	for i := 0; i < n; i++ {
		if curr == nil {
			return nil
		}
		curr = curr.Next
	}

	if curr == nil {
		return head.Next
	}

	for {
		prevOfDelay = delay
		delay = delay.Next
		curr = curr.Next
		if curr == nil {
			prevOfDelay.Next = delay.Next
			return head
		}
	}
}