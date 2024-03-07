package LinkedList

import (
	"fmt"
)

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (sll *SinglyLinkedList) Traverse() {
	current := sll.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
func (sll *SinglyLinkedList) Append(data interface{}) {
	newNode := &Node{Data: data}
	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = newNode
	} else {
		sll.Tail.Next = newNode
		sll.Tail = newNode
	}
	sll.Size++
}

func (sll *SinglyLinkedList) Prepend(data interface{}) {
	newNode := &Node{Data: data}
	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = newNode
	} else {
		newNode.Next = sll.Head
		sll.Head = newNode
	}
	sll.Size++
}
func (sll *SinglyLinkedList) InsertAtPosition(data interface{}, position int) {
	if position < 0 || position > sll.Size {
		println("Index out of bounds")
		return
	}
	newNode := &Node{Data: data}
	if position == 0 {
		sll.Prepend(data)
		return
	}
	if position == sll.Size {
		sll.Append(data)
		return
	}
	current := sll.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	sll.Size++
}
func (sll *SinglyLinkedList) DeleteFromHead() interface{} {
	if sll.Head == nil {
		return nil
	}
	deletedNode := sll.Head
	sll.Head = sll.Head.Next

	if sll.Head == nil {
		sll.Tail = nil
	}
	sll.Size--
	return deletedNode
}
func (sll *SinglyLinkedList) DeleteFromTail() interface{} {
	if sll.Head == nil {
		return nil
	}
	if sll.Size == 1 {
		deletedNode := sll.Head
		sll.Head = nil
		sll.Tail = nil
		sll.Size--
		return deletedNode
	}
	current := sll.Head
	for current.Next != sll.Tail {
		current = current.Next
	}
	deletedNode := sll.Tail
	sll.Tail = current
	sll.Tail.Next = nil
	sll.Size--
	return deletedNode
}

func (sll *SinglyLinkedList) DeleteAtPosition(position int) interface{} {
	if position < 0 || position > sll.Size {
		println("Index out of bounds")
		return nil
	}
	if position == 0 {
		return sll.DeleteFromHead()
	}
	if position == sll.Size {
		return sll.DeleteFromTail()
	}
	current := sll.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}
	deletedNode := current.Next
	current.Next = current.Next.Next
	sll.Size--
	return deletedNode
}

func (sll *SinglyLinkedList) FindMiddleNode() *Node {
	if sll.Head == nil {
		return nil
	}
	slow, fast := sll.Head, sll.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func (sll *SinglyLinkedList) Search(data interface{}) bool {
	current := sll.Head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}
func (sll *SinglyLinkedList) ReverseList() {
	var prev, curr, next *Node
	curr = sll.Head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	sll.Tail = sll.Head
	sll.Head = prev
}
func (sll *SinglyLinkedList) DisplayReverse() {
	if sll.Head == nil {
		return
	}
	current := sll.Head
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
