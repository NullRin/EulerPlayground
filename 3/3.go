package main

import (
	"math"
)

func simpleSolution(number int) int {
	var max = int(math.Sqrt(float64(number)))
	var factor = 2
	var result = 0
	for factor < max {
		if number%factor == 0 {
			result = number
			number = number / factor
			if number == 0 {
				break
			}
			factor = 2
			continue
		}
		factor++
	}
	return result
}

func main() {
	var number = 600851475143
	result := simpleSolution(number)
	print(result)
}
