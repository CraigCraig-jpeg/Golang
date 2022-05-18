package main 

import (
	"fmt"
	"context"
)

func main() {
	ctx := context.Background()
	// fmt.Printf("%v %T %v", a , v , ctx.Err())
	a , v := context.WithCancel(ctx)
	
	fmt.Printf("%v %T %v", a , v , ctx.Err())
}