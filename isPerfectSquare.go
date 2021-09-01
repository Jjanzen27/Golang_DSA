/*
/*
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Follow up: Do not use any built-in library function such as sqrt.



Example 1:

Input: num = 16
Output: true
Example 2:

Input: num = 14
Output: false


Constraints:

1 <= num <= 2^31 - 1


Input : number
Output : boolean

1 : 1
2 : 4
3 : 9
4 : 16
5 : 25

16 / 2 = 8 * 8 > 16 so 8 / 2 4 * 4 = 16

14 / 2 = 7 * 7 > 14


144 / 2 = 72 * 72 36



441 / 2 =220/ 110 / 55 / 27 / 13
/ 20 / 23 / ->21<-

input n
lastM
m =  n / 2
  if m square > n
   lastM = m
   m /=2
   else break
if m and lastM are only one apart, return false
now take the number halfway between m and lastM (m + ((lastM - m) / 2) call that p
m = p


*/

package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {

	candidate := num / 4
	previousCandidate := num

	if num < 4 {
		if num == 1 {
			return true
		} else {
			return false
		}
	}

	for previousCandidate-candidate > 1 {
		if candidate*candidate == num {
			return true
		} else if candidate*candidate > num {
			previousCandidate = candidate
			candidate = candidate / 2
		} else if candidate*candidate < num {
			gap := previousCandidate - candidate

			if gap%2 == 0 {
				gap = gap / 2
			} else {
				gap = (gap - 1) / 2
			}

			candidate = candidate + (gap)
			//fmt.Println(gap, previousCandidate, candidate*candidate == num)
		}

	}
	return candidate*candidate == num
}

func main() {
	fmt.Println(isPerfectSquare(36))
	fmt.Println(isPerfectSquare(25))
	fmt.Println(isPerfectSquare(49))
	fmt.Println(isPerfectSquare(64))
	fmt.Println(isPerfectSquare(100))

}
