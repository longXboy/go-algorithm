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
	ret := q.s2.pop()
	if ret != nil {
		return ret
	}
	for {
		d := q.s1.pop()
		if d == nil {
			break
		}
		q.s2.push(*d)
	}
	ret = q.s2.pop()
	return ret
}

func main() {
	var q Queue
	q.appendTail(1)
	q.appendTail(2)
	q.appendTail(3)
	fmt.Println(*q.deleteHead())
	fmt.Println(*q.deleteHead())
	q.appendTail(4)
	q.appendTail(5)
	fmt.Println(*q.deleteHead())
	fmt.Println(*q.deleteHead())
	fmt.Println(*q.deleteHead())
	q.deleteHead()
	q.appendTail(6)
	fmt.Println(*q.deleteHead())
	fmt.Println("===============")
	var s StackWithQ
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(*s.Pop())
	s.Push(4)
	s.Push(5)
	fmt.Println(*s.Pop())
	fmt.Println(*s.Pop())
	fmt.Println(*s.Pop())
	s.Push(6)
	fmt.Println(*s.Pop())
	fmt.Println(*s.Pop())

}

type StackWithQ struct {
	q1 Queue
	q2 Queue
}

func (s *StackWithQ) Push(v int) {
	s.q1.appendTail(v)
}

func (s *StackWithQ) Pop() *int {
	d := s.q1.deleteHead()
	if d != nil {
		for {
			newd := s.q1.deleteHead()
			if newd == nil {
				return d
				break
			}
			s.q2.appendTail(*d)
			d = newd
		}
	}
	d = s.q2.deleteHead()
	if d != nil {
		for {
			newd := s.q2.deleteHead()
			if newd == nil {
				return d
				break
			}
			s.q1.appendTail(*d)
			d = newd
		}
	}
	return nil
}
