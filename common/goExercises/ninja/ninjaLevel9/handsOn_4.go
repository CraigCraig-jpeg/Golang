package main 

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var increment int64
	var goRoutines int = 80
	var wg sync.WaitGroup
	wg.Add(goRoutines)
	
	for i := 0; i < goRoutines; i++ {
		go func (){
			atomic.AddInt64(&increment, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&increment))
			wg.Done()
		}()
	}
	wg.Wait()
}