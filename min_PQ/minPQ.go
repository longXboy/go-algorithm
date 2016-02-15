package main

import (
	"fmt"
)

func main() {
	quques := []int{11, 3, 100, -1, 16, 10, 5, 4, 7, 2, 21, 9, 66, 33, 1, 44, 3, 10, 12, 15, 99, 11, 34, 99}
	mSize := 16
	mpq := MinPQ{}
	mpq.Init(mSize)
	for _, data := range quques {
		mpq.Insert(data)
	}
	i := 1
	for {
		if i >= len(mpq.pq) {
			break
		}
		if i*2 < len(mpq.pq) {
			fmt.Println(mpq.pq[i : i*2])
		} else {
			fmt.Println(mpq.pq[i:])
		}
		i = i * 2
	}
}

type MinPQ struct {
	currentSize int
	pq          []int
}

func (mpq *MinPQ) Init(mSize int) {
	mpq.pq = make([]int, mSize+1)
	mpq.currentSize = 1
}
func (mpq *MinPQ) Insert(key int) {
	if mpq.currentSize >= len(mpq.pq) {
		mpq.DelMaxAndInsertToTop(key)
	} else {
		mpq.InsertToBottom(key)
		mpq.currentSize++
	}
}

func (mpq *MinPQ) DelMaxAndInsertToTop(key int) {
	if mpq.pq[1] <= key {
		return
	}
	mpq.pq[1] = key
	father := 1
	for {
		if father >= mpq.currentSize {
			break
		}
		child1 := father * 2
		child2 := father*2 + 1
		bigindex := 0
		if child1 <= mpq.currentSize-1 {
			bigindex = child1
		}
		if child2 <= mpq.currentSize-1 {
			if bigindex < 1 {
				bigindex = child2
			} else if mpq.pq[child2] > mpq.pq[bigindex] {
				bigindex = child2
			}

		}
		if bigindex < 1 {
			break
		}
		if mpq.pq[father] < mpq.pq[bigindex] {
			mpq.pq[father], mpq.pq[bigindex] = mpq.pq[bigindex], mpq.pq[father]
		} else {
			break
		}
		father = bigindex
	}
}

func (mpq *MinPQ) InsertToBottom(key int) {
	mpq.pq[mpq.currentSize] = key
	child := mpq.currentSize
	for {
		father := child / 2
		if father >= 1 {
			if mpq.pq[father] < mpq.pq[child] {
				mpq.pq[father], mpq.pq[child] = mpq.pq[child], mpq.pq[father]
			} else {
				break
			}
		} else {
			break
		}
		child = father
	}
}
