package main

import (
	"fmt"
	"math/rand"
)

const (
	P = 0.5
	MaxLevel = 8
)

func randomLevel() int {
	i := 1
	for i < MaxLevel {
		p := rand.Float64()
		if (p < P) {
			i++
		} else {
			break
		}
	}
	return i
}

type node struct {
	level int
	nexts []*node
	v     int
}

type skipList struct {
	head *node
}

func NewSkipList() *skipList {
	s := new(skipList)
	s.head = new(node)
	s.head.level = MaxLevel
	s.head.nexts = make([]*node, MaxLevel)
	return s
}

func (s *skipList) insert(val int) {
	l := randomLevel()
	addNode := new(node)
	addNode.level = l
	addNode.nexts = make([]*node, l)
	addNode.v = val
	i := l

	p := s.head
	for i > 0 {
		n := p.nexts[i - 1]
		for n != nil && n.v < val {
			p = n
			n = n.nexts[i - 1]
		}
		p.nexts[i - 1] = addNode
		addNode.nexts[i - 1] = n
		i --
	}
}

func (s *skipList) search(val int) *node {
	i := s.head.level
	p := s.head
	for i > 0 {
		n := p.nexts[i - 1]
		for n != nil && n.v < val {
			p = n
			n = n.nexts[i - 1]
		}
		if n != nil && n.v == val {
			return n
		}
		i --
	}
	return nil
}

func (s *skipList) delete(d *node) {
	i := d.level
	p := s.head
	for i > 0 {
		n := p.nexts[i - 1]
		for n != nil && n.v < d.v {
			p = n
			n = n.nexts[i - 1]
		}
		p.nexts[i - 1] = d.nexts[i - 1]
		i --
	}
}

func (s *skipList) print() {
	p := s.head
	for (p.nexts[0] != nil) {
		fmt.Println(p.nexts[0].v)
		p = p.nexts[0]
	}
}

func main() {

}
