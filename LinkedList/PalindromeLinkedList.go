/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	ptr1 := head
	ptr2 := head

	for ptr1 != nil && ptr1.Next != nil {
		ptr1 = ptr1.Next.Next
		ptr2 = ptr2.Next
	}

	prev := ptr2
	next := ptr2

	for ptr2 != nil {
		next = ptr2.Next
		ptr2.Next = prev

		prev = ptr2
		ptr2 = next
	}

	for head != prev {
		if head.Val != prev.Val {
			return false
		}

		head = head.Next
		prev = prev.Next
	}
	return true
}