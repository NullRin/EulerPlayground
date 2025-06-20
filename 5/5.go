package main

import "fmt"

func main() {
	var solution = 0
	var length = 20
	var i = length
	for {
		solution += 2
		i = length
		for i >= 3 {
			if solution%i != 0 {
				break
			}
			i--
		}
		if i == 2 {
			break
		}
	}
	fmt.Println(solution)

}
