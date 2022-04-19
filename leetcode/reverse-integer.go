//question https://leetcode.com/problems/reverse-integer/
//not completed
package main

import "fmt"

func main(){
	x := reverse(-1534236470)
	fmt.Println(x)
}
func reverse(x int) int{
	capu := 1534236469
	capd := x * -1
	if x >= capu {
		return 0
	}
	if x < 0 && capd > 1534236469{
		return 0
	}
	y := int32(x)
	if 0 > y {
		y = y * -1
		var rev_num int32
		for y > 0 {
			rev_num = rev_num*10 + y%10
			y = y/10
		}
	return int(rev_num * -1)
	}
	var rev_num int32
	for y > 0 {
		rev_num = rev_num*10 + y%10
		y = y/10
	}
	return int(rev_num)
}