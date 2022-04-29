// problem https://leetcode.com/problems/longest-common-prefix/
// Done 
package main 

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	rar := true
	for _, str := range strs {
		if str == "" {
			rar = false
		}
	}
	if len(strs) == 1 || rar == false{
		return strs[0]
	}else{
		var l string
		var bol bool = true
		var i int = 0
		for bol{
			x := strs[0]
			t := x[0:i]
			mon := 0
			for _ , v := range strs{
				if r := strings.HasPrefix(v, t) ; r == true{
					// fmt.Println(r , v , t) 
					mon ++ 
				}else {
					bol = false
					break
				}
			}
			if mon == len(strs){
				l = t
				mon = 0
			}
			if len(t) == len(strs){
				break
			}else{
				i++
			}
		}
		return l
	}
}

func main() {
	// input := []string{"flower","flow","flight"}
	// input := []string{""}

	// input := []string{"flower","flower","flower","flower"}
	x := longestCommonPrefix(input)
	fmt.Println(x)
	// fmt.Println(input)
	
}