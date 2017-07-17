package main

import (
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) pop() *int {
	if len(s.data) == 0 {
		return nil
	}
	d := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return &d
}

type Queue struct {
	s1 Stack
	s2 Stack
}

func (q *Queue) appendTail(i int) {
	q.s1.push(i)
}

func (q *Queue) deleteHead() *int {
	for {
		d := q.s1.pop()
		if d == nil {
			break
		}
		q.s2.push(*d)
	}
	ret := q.s2.pop()
	for {
		d := q.s2.pop()
		if d == nil {
			break
		}
		q.s1.push(*d)
	}
	return ret
}

func main() {
	var q Queue
	q.appendTail(1)
	q.appendTail(2)
	q.appendTail(3)
	fmt.Println(*q.deleteHead())
	q.appendTail(4)
	fmt.Println(*q.deleteHead())
	fmt.Println(*q.deleteHead())
	fmt.Println(*q.deleteHead())
	q.deleteHead()
	q.appendTail(5)
	fmt.Println(*q.deleteHead())
}
