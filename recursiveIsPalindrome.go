// Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Constraints:

// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	r := regexp.MustCompile("[^a-z0-9]")
	s = strings.ToLower(s)
	s = r.ReplaceAllString(s, "")
	return isPalindromeRecursive(s)
}

func isPalindromeRecursive(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	start, end := 0, len(s)-1

	if s[start] != s[end] {
		return false
	}

	return isPalindromeRecursive(s[start+1 : end])
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
