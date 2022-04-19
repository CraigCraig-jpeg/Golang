//question https://leetcode.com/problems/two-sum/
package main

import "fmt"

func main() {
	x := []int {3,2,3}
	y := 6
	z := twoSum(x , y)
	fmt.Println(z)
}
func twoSum(nums []int, target int) []int {
		x := nums
		y := target
		setter := []int{}
		for t , v := range x {
			for tt , vv := range x {
				if v + vv == y && t != tt{
						setter = []int{tt , t}
						break 
					}
				}
			}	
		return setter
	}