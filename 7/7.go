package main

import (
	"fmt"
	"math"
)

func main() {
	var limit = 10000000
	var primes = make([]bool, limit)
	var i = 2
	var j = i
	var index = 0
	var max = int(math.Sqrt(float64(limit))) + 1
	for i < max {
		j = i
		for {
			index = i * j
			if index >= limit {
				break
			}
			primes[index] = true
			j++
		}
		i++
	}
	i = -2
	for prime, notPrime := range primes {
		if !notPrime {
			i++
			// fmt.Println(prime)
			if i == 10001 {
				fmt.Println(prime)
			}
		}
	}
	fmt.Println(i)
}
