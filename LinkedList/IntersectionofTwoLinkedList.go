/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "math"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currA := headA
	currB := headB
	var lenA, lenB int
	for currA != nil || currB != nil {
		if currA != nil {
			lenA++
			currA = currA.Next
		}
		if currB != nil {
			lenB++
			currB = currB.Next
		}
	}

	diff := int(math.Abs(float64(lenA) - float64(lenB)))

	currA = headA
	currB = headB
	if lenA > lenB {
		for i := 0; i < diff; i++ {
			currA = currA.Next
		}
	} else {
		for i := 0; i < diff; i++ {
			currB = currB.Next
		}
	}

	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		} else {
			currA = currA.Next
			currB = currB.Next
		}
	}

	return nil
}