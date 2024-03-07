package LinkedList

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node // Used for doubly type linked list
}

// Interface for Common Linked List Operations
type LinkedListI interface {
	Append(data interface{})
	Prepend(data interface{})
	InsertAtPosition(data interface{}, position int)
	DeleteFromHead() interface{}
	DeleteFromTail() interface{}
	DeleteAtPosition(position int) interface{}
	Traverse()
	FindMiddleNode() *Node
	Search(data interface{}) bool
	ReverseList()
	DisplayReverse() []interface{}
}

// Interface for Circular Linked List Operations
type CircularLinkedListI interface {
	LinkedListI
	DetectLoop() bool
	FindIntersection(other CircularLinkedListI) *Node
}
