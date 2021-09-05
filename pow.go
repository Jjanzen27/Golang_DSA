// func main() {
// 	fmt.Println(myPow(2, 3))
// }

// package main

// import "fmt"

// func myPow(x float64, n int) float64 {
// 	if n == 1 {
// 		return x
// 	}

// 	return helper(x, n, x)
// }

// func helper(x float64, n int, a float64) float64 {
// 	if n == 2 {
// 		return a * x
// 	}

// 	a = a * x

// 	return helper(x, n-1, a)
// }

package main

import (
	"fmt"
	"math"
)

func myPow(x float64, n int) float64 {
	total := helper(x, int(math.Abs(float64(n))))

	if n >= 0 {
		return total
	} else {
		return 1 / total
	}
}

func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	total := helper(x*x, n/2)

	if n%2 != 0 {
		return x * total
	} else {
		return total
	}
}

func main() {
	fmt.Println(myPow(2, 9))
}
