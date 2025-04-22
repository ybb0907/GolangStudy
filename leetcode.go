package main

import (
	"fmt"
	"slices"
)

// func maxProfit1(prices []int) int {

// }

func rotate(nums []int, k int) {
	size := len(nums)
	k %= size
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func majorityElement(nums []int) int {
	res, fq := 0, 0
	for _, x := range nums {
		if res == x {
			fq++
		} else {
			fq--
			if fq == 0 {
				res = x
			}
		}
	}
	return res
}

func removeElement(nums []int, val int) int {
	k := 0
	for _, x := range nums {
		if x != val {
			nums[k] = x
			k++
		}
	}
	return k
	// i, j, size := 0, 0, len(nums)
	// for i < size {
	// 	if nums[i] != val {
	// 		nums[j] = nums[i]
	// 		i++
	// 		j++
	// 	} else {
	// 		i++
	// 	}
	// }
	// return j
}

func main() {
	fmt.Printf("111111\n")
}
