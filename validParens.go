package main

import "fmt"

type Stack struct {
	Top  *Node
	Size int
}

type Node struct {
	Value string
	Next  *Node
}

func (s *Stack) push(n *Node) {
	n.Next = s.Top
	s.Top = n
	s.Size++
}

func (s *Stack) pop() *Node {
	n := s.Top
	s.Top = s.Top.Next
	s.Size--
	return n
}

func (s *Stack) peek() string {
	return s.Top.Value
}

func isValid(s string) bool {
	stack := &Stack{}
	pairs := map[string]string{"(": ")", "{": "}", "[": "]"}

	for i := 0; i < len(s); i++ {
		char := string(s[i])

		if char == "(" || char == "[" || char == "{" {
			stack.push(&Node{Value: char})
		} else {
			if stack.Size > 0 && pairs[stack.peek()] == char {
				stack.pop()
			} else {
				return false
			}
		}
	}

	return stack.Size == 0
}

func main() {
	fmt.Println(isValid("]"))
	fmt.Println(isValid("((()"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("({{{)}}}"))
	fmt.Println(isValid("(}{}{}{)"))
	fmt.Println(isValid("{([])}"))
}
