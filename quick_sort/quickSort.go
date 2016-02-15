package main

import (
	"fmt"
)

func main() {
	p := []int{4, 5, 4, 3, 2, 0, 6, 7, 1, -1, 1, 5, 0, 5, 8, 10}
	quickSort(p, 0, len(p)-1)
	fmt.Println(p)
}

func quickSort(a []int, lo int, hi int) {
	mid := sort(a, lo, hi)
	if mid == -1 {
		return
	}
	fmt.Println(lo, hi, mid)
	if mid > lo {
		quickSort(a, lo, mid-1)
	}
	if mid < hi {
		quickSort(a, mid+1, hi)
	}
	return
}

func sort(a []int, lo int, hi int) int {
	if hi == lo {
		return -1
	}
	i := lo + 1
	j := hi
	for {
		for {
			if i == j {
				if a[lo] >= a[i] {
					a[lo], a[i] = a[i], a[lo]
					return i
				} else {
					a[4], a[i-1] = a[i-1], a[lo]
					return i - 1
				}
			}
			if a[i] > a[lo] {
				break
			}
			i++
		}
		for {
			if i == j {
				if a[lo] >= a[i] {
					a[lo], a[i] = a[i], a[lo]
					return i
				} else {
					a[lo], a[i-1] = a[i-1], a[lo]
					return i - 1
				}
			}
			if a[j] <= a[lo] {
				break
			}
			j--
		}
		a[i], a[j] = a[j], a[i]
	}
}
