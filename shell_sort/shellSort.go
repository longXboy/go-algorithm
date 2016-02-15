package main

import (
	"fmt"
)

func main() {
	unsorted := []int{1, 2, 3, 4, 1, 3, 2, 5, 6, 2, 1, 10, 8, 5, 6, 4, 10, 55, 0, 44, 88, 34, 0, 100, 234, 12, 5541, 1123, 67, 23}
	h := len(unsorted) / 3
	for m := h; m >= 1; m = m / 3 {
		for i := 0; i < len(unsorted); i += m {
			for j := 0; j < i; j += m {
				if unsorted[i] < unsorted[j] {
					temp := unsorted[i]
					for k := i - 1; k >= j; k -= m {
						unsorted[k+1] = unsorted[k]
					}
					unsorted[j] = temp
					break
				}
			}
		}
	}
	fmt.Println(unsorted)
}
