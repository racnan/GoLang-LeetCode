/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	head2 := head.Next
	ptr1 := head
	var ptr2 *ListNode

	isEven := false

	for ptr1.Next != nil {
		isEven = !isEven

		ptr2 = ptr1
		ptr1 = ptr1.Next

		ptr2.Next = ptr1.Next
	}

	if isEven {
		ptr2.Next = head2
	} else {
		ptr1.Next = head2
	}

	return head
}