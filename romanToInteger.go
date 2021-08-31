/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two one's added together.
12 is written as XII, which is simply X + II. The number 27 is written as XXVII,
which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII. Instead, the number four is written as IV.
Because the one is before the five we subtract it making four. The same principle applies to the number
nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.



Example 1:

Input: s = "III"
Output: 3
Example 2:

Input: s = "IV"
Output: 4
Example 3:

Input: s = "IX"
Output: 9
Example 4:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 5:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].


Problem: Convert string roman numeral to an integer
Input: A string comprised of the characters 'I', 'V', 'X', 'L', 'C', 'D', 'M',
    each representing a roman numeral
Output: An integer between 1 - 3999

Rules: don't have to validate input, roman letters will be capitalized
Three characters can impact the next character ('I', 'X', and 'C')
    Could check next char each time we encounter one of these or check previous once we hit
    one of the other characters

- Iterate through input string
- Each time we hit a mod-possible character, check prev character to see if mod necessary
- Add or subtract from total

*/
// --------------------------------------------------------------
package main

import (
	"fmt"
)

func rnToI(rn string) (total int) {
	var m = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i, v := range rn {

		if i == 0 {
			total += m[v]
			continue
		}

		if v == 'I' {
			total += 1
		} else if v == 'V' {
			if rn[i-1] == 'I' {
				total += 3
			} else {
				total += 5
			}
		} else if v == 'X' {
			if rn[i-1] == 'I' {
				total += 8
			} else {
				total += 10
			}
		} else if v == 'L' {
			if rn[i-1] == 'X' {
				total += 30
			} else {
				total += 50
			}
		} else if v == 'C' {
			if rn[i-1] == 'X' {
				total += 80
			} else {
				total += 100
			}
		} else if v == 'D' {
			if rn[i-1] == 'C' {
				total += 300
			} else {
				total += 500
			}
		} else if v == 'M' {
			if rn[i-1] == 'C' {
				total += 800
			} else {
				total += 1000
			}
		}

	}
	return
}

func main() {
	fmt.Println(rnToI("LVIII"))
	fmt.Println(rnToI("MCMXCIV"))
}

// _____________________________________________________________

// package main

// import (
// 	"fmt"
// )

func rnToI(rn string) (total int) {
	var m = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i, v := range rn {

		if i == 0 {
			total += m[v]
			continue
		}

		if v == 'V' || v == 'X' {
			if rn[i-1] == 'I' {
				total -= m['I'] + m[v]
			}
		} else if v == 'L' || v == 'C' {
			if rn[i-1] == 'X' {
				total -= m['X'] + m[v]
			}
		} else if v == 'D' || v == 'M' {
			if rn[i-1] == 'C' {
				total -= m['C'] + m[v]
			}
		} else {
			total += m[v]
		}
		fmt.Println(total)

	}
	return
}

func main() {
	fmt.Println(rnToI("LVIII"))
	fmt.Println(rnToI("MCMXCIV"))
}
