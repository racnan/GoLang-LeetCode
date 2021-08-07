/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	flattenMain(root)
	return root
}

func flattenMain(curr *Node) *Node {

	for {
		if curr.Next == nil {
			if curr.Child == nil {
				return curr
			} else {
				curr.Next = curr.Child
				curr.Child.Prev = curr
				curr.Child = nil
			}
		}
		if curr.Child != nil {
			next := curr.Next
			child := curr.Child

			curr.Next = curr.Child
			curr.Child.Prev = curr
			curr.Child = nil

			curr = flattenMain(child)

			curr.Next = next
			next.Prev = curr
			curr = next
		} else {
			curr = curr.Next
		}
	}
	return curr
}