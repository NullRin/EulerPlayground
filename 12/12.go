package main

import (
	"fmt"
	"math"
)

func main() {
	var triangleNumber = 0
	var i = 0
	var factor = 1
	var triangleRoot int
	var factorCount int16
	for factorCount < 500 {
		i++
		triangleNumber += i
		triangleRoot = int(math.Sqrt(float64(triangleNumber)))
		factorCount = 0
		factor = 2
		for factor <= triangleRoot {
			if triangleNumber%factor == 0 {
				factorCount += 2
				if factor == triangleRoot {
					factorCount--
				}
			}
			factor++
		}
	}
	fmt.Println(triangleNumber, factorCount)
}
