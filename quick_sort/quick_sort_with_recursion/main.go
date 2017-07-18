package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 1, 5, 10, 9, 6, 7, 8, 7, 11}
	QuickSort(a)
	fmt.Println(a)
	b := []byte("the quick fox jumps over a lazy dog")
	QuickSortBytes(b)
	fmt.Println(string(b))
}
func QuickSortBytes(a []byte) {
	if len(a) == 1 || len(a) == 0 {
		return
	}
	anchor := len(a) / 2
	anchorValue := a[anchor]
	a[anchor] = a[0]
	left := 0
	right := len(a) - 1
	for {
		if left == right {
			break
		}
		for ; right > left; right-- {
			if a[right] < anchorValue {
				a[left] = a[right]
				left++
				break
			}
		}
		for ; right > left; left++ {
			if a[left] >= anchorValue {
				a[right] = a[left]
				right--
				break
			}
		}
	}
	a[left] = anchorValue

	QuickSortBytes(a[:left])
	if left < (len(a) - 1) {
		QuickSortBytes(a[left+1:])
	}
}
func QuickSort(a []int) {
	if len(a) == 1 || len(a) == 0 {
		return
	}
	anchor := len(a) / 2
	anchorValue := a[anchor]
	a[anchor] = a[0]
	left := 0
	right := len(a) - 1
	for {
		if left == right {
			break
		}
		for ; right > left; right-- {
			if a[right] < anchorValue {
				a[left] = a[right]
				left++
				break
			}
		}
		for ; right > left; left++ {
			if a[left] >= anchorValue {
				a[right] = a[left]
				right--
				break
			}
		}
	}
	a[left] = anchorValue

	QuickSort(a[:left])
	if left < (len(a) - 1) {
		QuickSort(a[left+1:])
	}
}
