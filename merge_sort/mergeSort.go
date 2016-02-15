package main

import (
	"fmt"
	"math"
)

func main() {
	unsorted := []int{1, 2, 3, 4, 6666, 1, 3, 2, 5, 6, 2, 1, 10, 8, 5, 6, 4, 10, 55, 0, 44, 88, 34, 0, 100, 234, 12, 5541, 1123, 67, 23, 7, 12, 12, 123, 222}
	aux := make([]int, len(unsorted))
	m := Merge{}
	m.Sort(unsorted, aux)
	fmt.Println(unsorted)
	fmt.Println(aux)

}

type Merge struct {
}

func (m *Merge) less(a, b int) bool {
	if a < b {
		return true
	} else {
		return false
	}
}
func (m *Merge) Sort(un []int, aux []int) {
	m.sort(un, aux, 0, len(un)-1)
}
func (m *Merge) sort(un []int, aux []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	m.sort(un, aux, lo, mid)
	m.sort(un, aux, mid+1, hi)
	if int(math.Log2(float64(hi-lo+1)))%2 == 1 {
		m.merge(aux, un, lo, mid, hi)
	} else {
		m.merge(un, aux, lo, mid, hi)
	}

}

func (m *Merge) merge(un []int, aux []int, lo int, mid int, hi int) {
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			un[k] = aux[j]
			j++
		} else if j > hi {
			un[k] = aux[i]
			i++
		} else if m.less(aux[i], aux[j]) {
			un[k] = aux[i]
			i++
		} else {
			un[k] = aux[j]
			j++
		}
	}
}
