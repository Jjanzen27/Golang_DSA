// Given a string path, which is an absolute path (starting with a slash '/')
// to a file or directory in a Unix-style file system, convert it to the simplified canonical path.

// In a Unix-style file system, a period '.' refers to the current directory, a double period '..'
// refers to the directory up a level, and any multiple consecutive slashes (i.e. '//') are treated
// as a single slash '/'. For this problem, any other format of periods such as '...' are treated
// as file/directory names.

// The canonical path should have the following format:

// The path starts with a single slash '/'.
// Any two directories are separated by a single slash '/'.
// The path does not end with a trailing '/'.
// The path only contains the directories on the path from the root directory to the target file or
// directory (i.e., no period '.' or double period '..')
// Return the simplified canonical path.

// Example 1:

// Input: path = "/home/"
// Output: "/home"
// Explanation: Note that there is no trailing slash after the last directory name.
// Example 2:

// Input: path = "/../"
// Output: "/"
// Explanation: Going one level up from the root directory is a no-op, as the root level is the
// highest level you can go.
// Example 3:

// Input: path = "/home//foo/"
// Output: "/home/foo"
// Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
// Example 4:

// Input: path = "/a/./b/../../c/"
// Output: "/c"

// Constraints:

// 1 <= path.length <= 3000
// path consists of English letters, digits, period '.', slash '/' or '_'.
// path is a valid absolute Unix path.

/*

 */

package main

import (
	"fmt"
	"strings"
)

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

func (s *Stack) peek() *Node {
	return s.Top
}

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := &Stack{}

	for _, v := range parts {
		if v == "." || v == "" {
			continue
		} else if v == ".." {
			if stack.Size != 0 {
				stack.pop()
			} else {
				continue
			}
		} else {
			stack.push(&Node{Value: v})
		}
	}

	newPath := make([]string, stack.Size)

	for i := stack.Size - 1; stack.Size != 0; {
		val := stack.pop().Value
		newPath[i] = val
		i--
	}

	return "/" + strings.Join(newPath, "/")

}

func main() {
	fmt.Println(simplifyPath("/a/./b//../../c/"))
}
