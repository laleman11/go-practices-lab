package datastructures

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil}
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{next: nil, value: value}
}

func (l *LinkedList[T]) Add(value T) {
	newNode := NewNode(value)
	if l.head == nil {
		l.head = newNode
	} else {
		lastElement := l.head.findLast()
		lastElement.next = newNode
	}
}

func (n *Node[T]) findLast() *Node[T] {
	if n.next != nil {
		return n.next.findLast()
	} else {
		return n
	}
}

func (l *LinkedList[T]) Prepend(value T) {
	newNode := NewNode(value)
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList[T]) Delete(value T) {
	if l.head.value == value {
		l.head = l.head.next
	} else {
		node := l.head.next
		node.removeNode(l.head, value)
	}
}

func (n *Node[T]) removeNode(nodeBefore *Node[T], value T) {
	if n.value == value {
		nodeBefore.next = n.next
	} else {
		if n.next != nil {
			n.next.removeNode(n, value)
		}
	}
}

func (l *LinkedList[T]) Find(value T) (T, bool) {
	if l.head == nil {
		return value, false
	}
	return value, l.head.exist(value)
}

func (n *Node[T]) exist(value T) bool {
	if n.value == value {
		return true
	} else {
		if n.next != nil {
			return n.next.exist(value)
		} else {
			return n.value == value
		}
	}
}

func (l *LinkedList[T]) Size() int {
	if l.head == nil {
		return 0
	}
	node := l.head
	return node.size(0)
}

func (n *Node[T]) size(value int) int {
	if n.next == nil {
		return value + 1
	} else {
		return n.next.size(value + 1)
	}
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList[T]) Print() {
	if l.head != nil {
		l.head.print()
	} else {
		println("empty list")
	}
}

func (n *Node[T]) print() {
	if n.next == nil {
		println(n.value, "-> nil")
	} else {
		print(n.value, " -> ")
		n.next.print()
	}
}

func (l *LinkedList[T]) ToSlice() []T {
	if l.head == nil {
		return []T{}
	} else {
		return l.head.getSlice([]T{})
	}
}

func (n *Node[T]) getSlice(actualSlice []T) []T {
	if n.next == nil {
		return append(actualSlice, n.value)
	} else {
		actualSlice = append(actualSlice, n.value)
		return append(n.next.getSlice(actualSlice))
	}
}

func (l *LinkedList[T]) InsertAfter(value T, newValue T) {
	if l.head != nil {
		l.head.insertAfter(value, newValue)
	}
}

func (n *Node[T]) insertAfter(value T, newValue T) {
	if n.next == nil {
		if n.value == value {
			newNode := NewNode(newValue)
			n.next = newNode
		}
	} else {
		if n.value == value {
			newNode := NewNode(newValue)
			newNode.next = n.next
			n.next = newNode
		} else {
			n.next.insertAfter(value, newValue)
		}
	}
}

func (l *LinkedList[T]) Reverse() {
	l2 := NewLinkedList[T]()
	if l.head != nil {
		actualNode := l.head.removeLast(nil)
		l2.head = actualNode
		if l.head.next != nil {
			for l.head != nil {
				actualNode.next = l.head.removeLast(actualNode)
				actualNode = actualNode.next
				if actualNode == l.head {
					l.head = nil
				}
			}
		}
	}

	l.head = l2.head
}

func (n *Node[T]) removeLast(nodeBefore *Node[T]) *Node[T] {
	if n.next == nil {
		response := n
		if nodeBefore != nil {
			nodeBefore.next = nil
		}
		return response
	} else {
		return n.next.removeLast(n)
	}
}
