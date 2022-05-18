//question https://leetcode.com/problems/add-two-numbers/
//Done

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     next *ListNode
//  * }
//  */

package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	next *ListNode
	}

func addTwoNumbers(l1 *ListNode)  {
	n := []int{} 
	var f bool = true
	n = append(n, l1.Val)
	current := l1
	for f {
		if current.next == nil {
			f = false
			break
		}
		current = current.next
		n = append(n, current.Val)
	}
	// fmt.Println(n)
}



func main() {
	
	node0 := &ListNode{
		Val: 2,
	}
	node1 := &ListNode{
		Val: 4,
	}
	// node2 := &ListNode{
	// 	Val: 3,
	// }

	// node0B := &ListNode{
	// 	Val: 5,
	// }
	// node1B := &ListNode{
	// 	Val: 6,
	// }
	// node2B := &ListNode{
	// 	Val: 4,
	// }
	
	addTwoNumbers(node0)
	fmt.Println(node1)
}