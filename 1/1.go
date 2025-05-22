package main

import "fmt"

func main() {
	var limit = 1000
	var mult1 = 3
	var mult2 = 5
	var x = 0
	var i = 1
	var lastMult = i * mult1
	for lastMult < limit {
		x += lastMult
		i++
		lastMult = i * mult1
	}
	i = 1
	lastMult = i * mult2
	for lastMult < limit {
		x += lastMult
		i++
		lastMult = i * mult2
	}
	var multDup = mult1 * mult2
	i = 1
	lastMult = i * multDup
	for lastMult < limit {
		x -= lastMult
		i++
		lastMult = i * multDup
	}
	fmt.Println(x)
}
