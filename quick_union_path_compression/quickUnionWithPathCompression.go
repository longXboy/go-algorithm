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
	qu := QuickUnionWithPathCompression{}
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
	stack := TempStack{}
	stack.init(500)
	start := time.Now().UnixNano()
	for i := range inputNums {
		qu.Union(inputNums[i].N1, inputNums[i].N2, &stack)
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)
	fmt.Println(qu.Count())
}

type QuickUnionWithPathCompression struct {
	id    []int
	sz    []int
	count int
}

func (qu *QuickUnionWithPathCompression) Init(n int) {
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

func (qu *QuickUnionWithPathCompression) Count() int {
	return qu.count
}

func (qu *QuickUnionWithPathCompression) Connected(p int, q int, stack *TempStack) bool {
	return qu.find(p, stack) == qu.find(q, stack)
}

func (qu *QuickUnionWithPathCompression) find(p int, stack *TempStack) int {
	n := 0
	for {
		if qu.id[p] == p {
			for i := 0; i < n; i++ {
				qu.id[stack.a[i]] = p
			}
			return p
		}
		//	oldp := p
		if n == stack.count() {
			stack.resize(n * 2)
		}
		p = qu.id[p]
		stack.a[n] = p
		n++
		//qu.id[oldp] = qu.id[p]
	}
}

func (qu *QuickUnionWithPathCompression) Union(p int, q int, stack *TempStack) {
	i := qu.find(p, stack)
	j := qu.find(q, stack)
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

type TempStack struct {
	a    []int
	size int
}

func (s *TempStack) init(size int) {
	s.a = make([]int, size)
	s.size = size
}
func (s *TempStack) count() int {
	return s.size
}

func (s *TempStack) resize(newsize int) {
	newa := make([]int, newsize)
	for i := range s.a {
		fmt.Println("asd")
		newa[i] = s.a[i]
	}
	s.a = newa
}
