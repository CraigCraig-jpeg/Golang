package main 

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var increment int 
	var goRoutines int = 80
	var wg sync.WaitGroup
	wg.Add(goRoutines)
	var mu sync.Mutex
	
	for i := 0; i < goRoutines; i++ {
		go func (){
			mu.Lock()
			v := increment
			runtime.Gosched()
			v ++ 
			increment = v
			fmt.Println(increment)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}