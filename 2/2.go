package main

import "fmt"

func main() {
	var max = 4000000
	var l1 = 1
	var l2 = 2
	var temp = 0
	var sum = 0
	for l2 < max {
		if l2&1 == 0 {
			sum += l2
		}
		temp = l2
		l2 += l1
		l1 = temp
	}
	fmt.Println(sum)
}
