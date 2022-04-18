// question https://codeforces.com/problemset/problem/746/B
package main 

import "fmt"
import "math"

func main(){
	word := []string{"l" , "o" , "g" , "v" , "a"}
	fmt.Println(word)
	for i := 0 ; i < 4 ; i++{
	//find median
	middle := math.Ceil(float64(len(word))/2.0)
	if (len(word) % 2 == 0){
		middle = float64((len(word)/2)-1)
		input := word[int(middle)]
		word = append(word[:int(middle)] ,word[int(middle)+1:]...)
		word = append([]string{input} , word...)
		fmt.Println(word)
	}else{
	input := word[int(middle)-1]
	word = append(word[:int(middle)-1] ,word[int(middle):]...)
	word = append([]string{input} , word...)
	fmt.Println(word)
	}
	}
	//replace median with
}