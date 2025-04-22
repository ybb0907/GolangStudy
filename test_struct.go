package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

// fibonacci 是返回一个「返回一个 int 的函数」的函数
func fibonacci() func() int {
	s1 := 0
	s2 := 1

	return func() int {
		sum := s1 + s2
		s1 = s2
		s2 = sum
		return sum
	}
}

func WordCount(s string) map[string]int {

	res_map := make(map[string]int)

	res := make([]string, 0, 100)
	res = strings.Fields(s)

	for _, v := range res {
		if _, ok := res_map[v]; ok {
			res_map[v]++
		} else {
			res_map[v] = 1
		}
	}

	return res_map
}

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz  foo  "))
}
