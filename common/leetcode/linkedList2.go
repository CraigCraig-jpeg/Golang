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

func (l linkedNode) printList() {
	toPrint := l.head 
	for l.length != 0 {
		fmt.Println(toPrint.data)
		toPrint = toPrint.next
		l.length--
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
		data: 10,
	}
	list.changeMe(num2)
	list.printList()
	
}