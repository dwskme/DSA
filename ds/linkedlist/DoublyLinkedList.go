package LinkedList

import (
	"fmt"
)

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (dll *DoublyLinkedList) Traverse() {
	current := dll.Head
	for current != nil {
		fmt.Printf("%v <-> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (dll *DoublyLinkedList) Append(data interface{}) {
	newNode := &Node{Data: data}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		dll.Tail.Next = newNode
		newNode.Prev = dll.Tail
		dll.Tail = newNode
	}
	dll.Size++
}

func (dll *DoublyLinkedList) Prepend(data interface{}) {
	newNode := &Node{Data: data}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
	dll.Size++
}
func (dll *DoublyLinkedList) InsertAtPosition(data interface{}, position int) {
	if position < 0 || position > dll.Size {
		println("Index out of bounds")
		return
	}
	newNode := &Node{Data: data}
	if position == 0 {
		dll.Prepend(data)
		return
	}
	if position == dll.Size {
		dll.Append(data)
		return
	}
	current := dll.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	newNode.Prev = current
	current.Next.Prev = newNode
	current.Next = newNode
	dll.Size++
}
func (dll *DoublyLinkedList) DeleteFromHead() interface{} {
	if dll.Head == nil {
		return nil
	}
	deletedNode := dll.Head
	dll.Head = dll.Head.Next

	if dll.Head == nil {
		dll.Tail = nil
	}
	dll.Size--
	return deletedNode
}
func (dll *DoublyLinkedList) DeleteFromTail() interface{} {
	if dll.Head == nil {
		return nil
	}
	if dll.Size == 1 {
		deletedNode := dll.Head
		dll.Head = nil
		dll.Tail = nil
		dll.Size--
		return deletedNode
	}
	current := dll.Head
	for current.Next != dll.Tail {
		current = current.Next
	}
	deletedNode := dll.Tail
	dll.Tail = current
	dll.Tail.Next = nil
	dll.Size--
	return deletedNode
}

func (dll *DoublyLinkedList) DeleteAtPosition(position int) interface{} {
	if position < 0 || position > dll.Size {
		println("Index out of bounds")
		return nil
	}
	if position == 0 {
		return dll.DeleteFromHead()
	}
	if position == dll.Size {
		return dll.DeleteFromTail()
	}
	current := dll.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}
	deletedNode := current.Next
	current.Next = current.Next.Next
	dll.Size--
	return deletedNode
}

func (dll *DoublyLinkedList) FindMiddleNode() *Node {
	if dll.Head == nil {
		return nil
	}
	slow, fast := dll.Head, dll.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func (dll *DoublyLinkedList) Search(data interface{}) bool {
	current := dll.Head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}
func (dll *DoublyLinkedList) ReverseList() {
	var prev, curr, next *Node
	curr = dll.Head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		curr.Prev = next
		prev = curr
		curr = next
	}
	dll.Tail = dll.Head
	dll.Head = prev
}
func (dll *DoublyLinkedList) DisplayReverse() {
	if dll.Head == nil {
		return
	}
	current := dll.Head
	var printReverse func(*Node)
	printReverse = func(node *Node) {
		if node.Next != nil {
			printReverse(node.Next)
		}
		fmt.Print(node.Data, " ")
	}
	printReverse(current)
	fmt.Println()
}
