//question https://leetcode.com/problems/two-sum/
package main

import "fmt"

func main() {
	x := []int {-1,-2,-3,-4,-5}
	y := -8
	z := twoSum(x , y)
	fmt.Println(z)
}
func twoSum(nums []int, target int) []int {
		x := nums
		y := target
		counter := 0
		setter := []int{}
		
		for t , v := range x {
			for tt , vv := range x {
				if v + vv == y {
					counter += 1
					if counter == 2 {
						setter = []int{tt , t}
						break 
					}
				}
			}
		}
		return setter
	}