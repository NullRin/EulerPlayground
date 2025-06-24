package main

import (
	"fmt"
	"math"
)

func main() {
	var limit = 2000000
	var notPrimes = make([]bool, limit)
	notPrimes[0] = true
	notPrimes[1] = true
	var i = 1
	var j = i
	var index = 0
	var max = int(math.Sqrt(float64(limit))) + 1
	for i < max {
		i++
		j = i
		if notPrimes[i] {
			continue
		}
		for {
			index = i * j
			if index >= limit {
				break
			}
			notPrimes[index] = true
			j++
		}
	}
	var sumOfPrimes = 0
	for prime, notPrime := range notPrimes {
		if !notPrime {
			sumOfPrimes += prime
			// fmt.Println(prime)
		}
	}
	fmt.Println(sumOfPrimes)
}
