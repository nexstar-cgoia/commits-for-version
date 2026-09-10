package linkedlist

import "iter"

type Node struct {
	SHA     *string
	Message *string
	Next    *Node
}

type AppendNode struct {
	SHA     *string
	Message *string
}

type iterNode struct {
	SHA     *string
	Message *string
}

// For this use case i only want a reference to the head ( which will be the oldest commit ) and from there to iterate through the linked list.
type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(n AppendNode) {
	node := &Node{SHA: n.SHA, Message: n.Message}
	if ll.Head == nil {
		ll.Head = node
	} else {
		node.Next = ll.Head
		ll.Head = node
	}
}

func (ll *LinkedList) Iter() iter.Seq[iterNode] {
	return func(yield func(iterNode) bool) {
		for current := ll.Head; current != nil; current = current.Next {
			if !yield(iterNode{SHA: current.SHA, Message: current.Message}) {
				break
			}
		}
	}
}
