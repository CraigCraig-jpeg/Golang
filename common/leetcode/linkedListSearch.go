package main

import (
	"fmt"
)

type node struct {
	data int
	next *node 
}

type linkedNode struct {
	head *node
	length int 
}

func (l *linkedNode) changeMe(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedNode) browse() {
	first := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println(first.data)
		first = l.head.next
	}
}

func main() {
	list := linkedNode{

	}

	num := &node{
		data: 10,
	}

	list.changeMe(num)
	num2 := &node{
		data: 11,
	}
	list.changeMe(num2)
	list.browse()
}