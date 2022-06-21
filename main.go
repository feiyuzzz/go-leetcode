package main

import "github.com/feiyuzzz/go-leetcode/sort"

func main() {
	a := []int{3, 1, 4}
	sort.QuickSort(a)
	for i := 0; i < len(a); i++ {
		print(a[i])
	}
}
