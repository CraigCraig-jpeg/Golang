package main

import (
	"fmt"
)

type person struct {
	name string 
	next *person
}

type agents struct {
	head *person
	length int
}

func (a *agents) preappend(n *person) {
	second := a.head
	a.head = n 
	a.head.next = second
	a.length++
}

func main() {
	node := &person {
		name: "Craig",
	}
	node1 := &person {
		name: "Jack",
	}
	list := agents{

	}
	list.preappend(node)
	list.preappend(node1)

	fmt.Println(list)
}

