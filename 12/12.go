package main

import (
	"fmt"
	"math"
)

func main() {
	var factorGoal int16 = 500
	var triangleNumber = 0
	var i = 0
	var factor = 1
	var triangleRoot int
	var factorCount int16
	for factorCount < factorGoal {
		i++
		triangleNumber += i
		triangleRoot = int(math.Sqrt(float64(triangleNumber)))
		factorCount = 2
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
