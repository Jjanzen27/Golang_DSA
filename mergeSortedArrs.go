// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside
// the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements
// denote the elements that should be merged, and the last n elements are set to 0 and should be
// ignored. nums2 has a length of n.

// Example 1:

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
// Example 2:

// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The result of the merge is [1].
// Example 3:

// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
// Output: [1]
// Explanation: The arrays we are merging are [] and [1].
// The result of the merge is [1].
// Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the
// merge result can fit in nums1.

// Constraints:

// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -109 <= nums1[i], nums2[j] <= 109

// Follow up: Can you come up with an algorithm that runs in O(m + n) time?

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	m--
	n--
	i := len(nums1) - 1

	for m >= 0 && n >= 0 {

		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}

	for m >= 0 {
		nums1[i] = nums1[m]
		m--
		i--
	}

	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
	return
}

func main() {
	// 	arr1 := []int{1, 2, 3, 0, 0, 0}
	// 	arr2 := []int{2, 5, 6}
	// 	arr3 := []int{0}
	// 	arr4 := []int{1}
	arr5 := []int{2, 0}
	arr6 := []int{1}
	// 	merge(arr1, 3, arr2, 3)
	// 	merge(arr3, 0, arr4, 1)
	merge(arr5, 1, arr6, 1)
	// 	fmt.Println(arr1)
	// 	fmt.Println(arr3)
	fmt.Println(arr5)
}
