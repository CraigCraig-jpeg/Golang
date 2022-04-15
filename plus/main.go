package main 

import "fmt" 

type hotdog int
var z hotdog

func main() {
    x := 50
    x = 50 + 2
    var y int = 43
    z = hotdog(x)

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
}