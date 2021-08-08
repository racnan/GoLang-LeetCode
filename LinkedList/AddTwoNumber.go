//Runtime: 8ms; Runtime Percentile: 88.06%
//Memory: 5MB; Memory Percentile: 94.35%

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carr int
	ptr := &ListNode{}
	ans := ptr
	for {
		ptr.Val = l1.Val + l2.Val + carr
		carr = 0
		if ptr.Val >= 10 {
			ptr.Val = ptr.Val % 10
			carr = 1
		}
		if l1.Next == nil && l2.Next == nil && carr == 0 {
			return ans
		}
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}
		ptr.Next = &ListNode{}
		ptr = ptr.Next

	}
}
