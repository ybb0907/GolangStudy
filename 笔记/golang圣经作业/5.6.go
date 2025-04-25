package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topo(courses map[string][]string) []string {
	used := make(map[string]bool)
	order := []string{}
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, val := range items {
			if !used[val] {
				used[val] = true
				visitAll(courses[val])
				order = append(order, val)
			}
		}
	}

	var keys []string
	for key := range courses {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	fmt.Println(topo(prereqs))
}
