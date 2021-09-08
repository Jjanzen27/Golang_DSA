package main

import "fmt"

type Stack struct {
	Top  *Node
	Size int
}

type Node struct {
	Value interface{} // All types satisfy the empty interface, so we can store anything here.
	Next  *Node
}

func main() {
	stack := &Stack{}
	node1 := &Node{Value: "horse"}
	node2 := &Node{Value: "sheep"}
	node3 := &Node{Value: "cow"}

	stack.push(node1)
	fmt.Println(stack.peek())
	stack.pop()
	fmt.Println(stack.peek())
	stack.push(node2)
	stack.push(node3)
	fmt.Println(stack.peek())
	stack.pop()
	fmt.Println(stack.peek())
	return
}

func (s *Stack) push(n *Node) {
	n.Next = s.Top
	s.Top = n
	return
}

func (s *Stack) pop() *Node {
	n := s.Top
	s.Top = s.Top.Next
	return n
}

func (s *Stack) peek() *Node {
	return s.Top
}
