/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	curr := head

	for curr != nil {
		curr.Next = &Node{Val: curr.Val, Next: curr.Next}
		curr = curr.Next.Next
	}

	curr = head

	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		} else {
			curr.Next.Random = nil
		}

		curr = curr.Next.Next
	}

	curr = head

	head = curr.Next
	copied := curr.Next

	for {
		curr.Next = curr.Next.Next
		curr = curr.Next

		if curr != nil {
			copied.Next = curr.Next
			copied = copied.Next
		} else {
			copied.Next = nil
			break
		}
	}

	return head
}