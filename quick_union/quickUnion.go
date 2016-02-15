package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type InputNum struct {
	N1 int
	N2 int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: quick_union [filename]")
		return
	}
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("read file(%s) failed!err :=%v", os.Args[1], err)
		return
	}
	slice := strings.Split(string(content), "\n")
	num, err := strconv.ParseInt(slice[0], 10, 32)
	if err != nil {
		fmt.Println("convert slice[0] to int failed!err:=%v ", err)
		return
	}
	qu := WeightedQuickUnionUF{}
	qu.Init(int(num))
	inputNums := make([]InputNum, 0)
	for _, data := range slice[1:] {
		if len(data) < 2 {
			continue
		}
		nums := strings.Split(data, " ")
		n1, err := strconv.ParseInt(nums[0], 10, 32)
		if err != nil {
			fmt.Println("convert nums[0] to int failed!err:=%v ", err)
			return
		}
		n2, err := strconv.ParseInt(nums[1], 10, 32)
		if err != nil {
			fmt.Println("convert nums[1] to int failed!err:=%v ", err)
			return
		}
		inputNums = append(inputNums, InputNum{N1: int(n1), N2: int(n2)})
	}
	start := time.Now().UnixNano()
	for i := range inputNums {
		qu.Union(inputNums[i].N1, inputNums[i].N2)
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)
	fmt.Println(qu.Count())
}

type WeightedQuickUnionUF struct {
	id    []int
	sz    []int
	count int
}

func (qu *WeightedQuickUnionUF) Init(n int) {
	qu.count = n
	qu.id = make([]int, n)
	for i := range qu.id {
		qu.id[i] = i
	}
	qu.sz = make([]int, n)
	for i := range qu.sz {
		qu.sz[i] = 1
	}
}

func (qu *WeightedQuickUnionUF) Count() int {
	return qu.count
}

func (qu *WeightedQuickUnionUF) Connected(p int, q int) bool {
	return qu.find(p) == qu.find(q)
}

func (qu *WeightedQuickUnionUF) find(p int) int {
	for {
		if qu.id[p] == p {
			return p
		}
		p = qu.id[p]
	}
}

func (qu *WeightedQuickUnionUF) Union(p int, q int) {
	i := qu.find(p)
	j := qu.find(q)
	if i == j {
		return
	}
	if qu.sz[i] < qu.sz[j] {
		qu.id[i] = qu.id[j]
		qu.sz[j] += qu.sz[i]
	} else {
		qu.id[j] = qu.id[i]
		qu.sz[i] += qu.sz[j]
	}
	qu.count--
}
