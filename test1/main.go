package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	res := make([]int, len(nums1)+len(nums2))
	if len(nums2) == 0 {
		return
	}
	n1, n2 := 0, 0
	for n1 < len(nums1) && n2 < len(nums2) {
		if nums1[n1] > nums2[n2] {
			res = append(res, nums1[n1])
			n1++
		} else {
			res = append(res, nums2[n2])
			n2++
		}
	}
	for n1 < len(nums1) {
		res = append(res, nums1[n1])
		n1++
	}

	for n2 < len(nums1) {
		res = append(res, nums2[n2])
		n2++
	}

	nums1 = res
}

func main() {
	var name, flag, age = "ybb", false, 25
	fmt.Printf("name: %v\n", name)
	fmt.Printf("flag: %v\n", flag)
	fmt.Printf("age: %v\n", age)

	const (
		b = iota
	)

	fmt.Printf("%T\n", b)

	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T\n", a)
}
