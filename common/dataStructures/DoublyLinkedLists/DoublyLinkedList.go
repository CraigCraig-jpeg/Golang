// package https://pkg.go.dev/container/list@go1.18.1

package main

import (
	"container/list"
	"fmt"
)

func main() {
	values := list.New()
	values.PushBack("Bird")
	values.PushBack("cat")
	values.PushFront("Snake")
	values.

	for e := values.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}