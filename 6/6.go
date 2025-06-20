package main

import "fmt"

func main() {
	var n = 100
	var sqrSum = (n * (n + 1) * (2*n + 1)) / 6
	var sum = (n * (n + 1)) / 2
	var solution = (sum * sum) - sqrSum
	fmt.Println(solution)
}
