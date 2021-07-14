type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	length int
	head   *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{0, nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {

	if index > this.length || index < 0 {
		return -1
	}

	curr := this.head
	currIndex := 0

	for curr != nil {
		if currIndex == index {
			return curr.val
		} else {
			curr = curr.next
			currIndex++
		}
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {

	newNode := Node{val, this.head}
	this.head = &newNode
	this.length++
	return
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {

	newNode := Node{val, nil}

	curr := this.head

	if curr == nil {
		newNode.next = this.head
		this.head = &newNode
		this.length++
		return
	}

	for {
		if curr.next == nil {
			curr.next = &newNode
			this.length++
			return
		} else {
			curr = curr.next
		}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index > this.length || index < 0 {
		return
	}

	newNode := Node{val, nil}

	if index == 0 {
		newNode.next = this.head
		this.head = &newNode
		this.length++
		return
	}

	curr := this.head
	prev := this.head
	currIndex := 0
	for {
		if currIndex == index {
			newNode.next = curr
			prev.next = &newNode
			this.length++
			return
		} else if curr == nil {
			return
		} else {
			prev = curr
			curr = curr.next
			currIndex++
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {

	if index > this.length || index < 0 {
		return
	}

	if index == 0 {
		curr := this.head
		this.head = curr.next
		this.length--
		return
	}

	curr := this.head
	prev := this.head
	currIndex := 0
	for {
		if currIndex == index {
			prev.next = curr.next
			this.length--
			return
		} else if curr.next == nil {
			return
		} else {
			prev = curr
			curr = curr.next
			currIndex++
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */