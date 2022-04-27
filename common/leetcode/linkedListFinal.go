//question https://leetcode.com/problems/add-two-numbers/
//Done
package main

import (
	"fmt"
	"strconv"
	"strings"
)


  	// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

type ListNodeList struct {
	head *ListNode
	length int
}

func (l *ListNodeList) prepend(node *ListNode){
	second := l.head 
	l.head = node 
	l.head.Next = second
	l.length++
}

func (l ListNodeList) total() []int {
	swap := l.head
	var t []int
	// t := make([]int , l.length , l.length)
	for i := 0; i < l.length; i++ {
		t = append(t, swap.Val)
		swap = swap.Next
	}
	return t
}  

func val(x []int) int{
	values := x
	valuesText := []string{}
	for i := range values {
        number := values[i]
        text := strconv.Itoa(number)
        valuesText = append(valuesText, text)
    }
	result := strings.Join(valuesText, "")
	intVar, _ := strconv.Atoi(result)
	return intVar
}
//  func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     return l1
// }

func main() {
	node0 := &ListNode{
		Val: 1,
	}
	node1 := &ListNode{
		Val: 2,
	}
	node2 := &ListNode{
		Val: 3,
	}
	node3 := &ListNode{
		Val: 3,
	}
	list := &ListNodeList{

	}
	list.prepend(node0)
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)

	val1 := val(list.total())
	
	node0B := &ListNode{
		Val: 4,
	}
	node1B := &ListNode{
		Val: 4,
	}
	node2B := &ListNode{
		Val: 4,
	}
	node3B := &ListNode{
		Val: 4,
	}
	listB := &ListNodeList{

	}
	listB.prepend(node0B)
	listB.prepend(node1B)
	listB.prepend(node2B)
	listB.prepend(node3B)

	val2 := val(listB.total())

	fmt.Println(val1 + val2)
}	

