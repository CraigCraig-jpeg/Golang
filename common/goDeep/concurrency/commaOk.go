package main
import (
	"fmt"
)

func main() {
		even := make(chan int)
		odd := make(chan int)
		quit := make(chan bool)

		go send(even, odd, quit)

		receive(even, odd, quit)

		fmt.Println("about to exit")
	}

	// send channel
	func send(even, odd chan<- int, quit chan<- bool) {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(quit)
	}

	// receive channel
	func receive(even, odd <-chan int, quit <-chan bool) {
		for {
			select {
			case v := <-even:
				fmt.Println("the value received from the even channel:", v)
			case v := <-odd:
				fmt.Println("the value received from the odd channel:", v)
			case i, ok := <-quit:
				if !ok {
					fmt.Println("from comma ok bit", i)
					return
				} else {
					fmt.Println("from comma ok bit", i)
				}
				/* this is a odd case i will explain why...
				odd <- i means it assignes the value to the odd channel
				x := <-odd means it pulls the value from the odd channel and assignes it
				to the value x 
				x , ok := <- odd means it pulls the value and assines it to the value.
				the value ok is a bool that becomes fulse once the channel is closed and 
				there are no more values to be polled off the odd channel
				find https://go.dev/tour/concurrency/4 for more information
				*/
			}
		}
	}
