package linkedlist

import "fmt"

type Node struct {
	Val  interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func Init() *LinkedList {
	head := &Node{nil, nil}
	list := &LinkedList{head}
	return list
}

func (ll *LinkedList) Fill(vals []int) {
	for _, v := range vals {
		ll.Add(v)
	}
}

func (ll *LinkedList) GetIndex(index int) *Node {
	indirect := ll.head
	for i := 0; i < index; i++ {
		indirect = indirect.next
	}

	return indirect
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%+v->", current.Val)

		current = current.next
	}

	fmt.Printf("%s \n", "nil")
	//fmt.Println(ll.head.Val)
}

func (ll *LinkedList) Add(Val interface{}) {

	if ll.head.Val == nil {
		ll.head = &Node{Val, nil}
	} else {

		current := ll.head

		for current.next != nil {
			current = current.next
		}
		current.next = &Node{Val, nil}
	}

}

func (ll *LinkedList) Remove(entry *Node) {
	indirect := ll.head

	for indirect.next != entry {
		indirect = indirect.next
		//fmt.Print("moving")

	}

	// then just remove it
	indirect.next = entry.next

}
