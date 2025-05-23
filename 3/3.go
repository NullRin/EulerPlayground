package main

import (
	"fmt"
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
			if number == 1 {
				break
			}
			factor = 2
			continue
		}
		factor++
	}
	return result
}

func nextPrime(primes []int) int {
	value := primes[len(primes)-1]
	prime := 0
	isPrime := false
	for {
		isPrime = true
		for _, prime = range primes {
			if value%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			break
		}
		value += 2
	}
	// fmt.Println(value)
	return value
}

func primeSolution(number int) int {
	var max = int(math.Sqrt(float64(number)))
	var result = 0
	var primes = []int{2, 3, 5, 7, 11, 13}
	var factor = primes[0]
	var primeCount = 0
	var temp float64 = 0
	var remainder float64 = 0
	var tempFloor float64 = 0
	for factor < max {
		temp = float64(number) / float64(factor)
		tempFloor = float64(int(temp))
		remainder = (temp - tempFloor) * tempFloor
		if remainder == 0 {
			result = number
			number = int(temp)
			if number == 1 {
				break
			}
			continue
		}
		primeCount++
		if len(primes) < primeCount {
			primes = append(primes, nextPrime(primes))
		}
		factor = primes[primeCount-1]
	}
	return result
}

func primeFloorSolution(number int) int {
	var max = int(math.Sqrt(float64(number)))
	var result = 0
	var primes = []int{2, 3, 5, 7, 11, 13}
	var factor = primes[0]
	var primeCount = 0
	var temp float64 = 0
	var remainder float64 = 0
	var tempFloor float64 = 0
	for factor < max {
		temp = float64(number) / float64(factor)
		tempFloor = float64(int(temp))
		remainder = (temp - tempFloor) * tempFloor
		if remainder == 0 {
			result = number
			number = int(temp)
			if number == 1 {
				break
			}
			continue
		}
		primeCount++
		if len(primes) < primeCount {
			primes = append(primes, nextPrime(primes))
		}
		factor = primes[primeCount-1]
	}
	return result
}

func main() {
	var number = 600851475143
	result := 0
	// result = simpleSolution(number)
	// fmt.Println(result)
	// result = 0
	// result = primeSolution(number)
	result = 0
	result = primeFloorSolution(number)
	fmt.Println(result)
}
