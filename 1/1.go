package main

import "fmt"

func main() {
	var limit = 1000
	var mult1 = 3
	var mult2 = 5
	var x = 0
	var i = 1
	var lastMult = 0
	for lastMult <= limit-mult1 {
		lastMult = i * mult1
		x += lastMult
		i++
	}
	i = 1
	lastMult = 0
	for lastMult <= limit-mult2 {
		lastMult = i * mult2
		x += lastMult
		i++
	}
	var multDup = mult1 * mult2
	i = 1
	lastMult = 0
	for lastMult <= limit-multDup {
		lastMult = i * multDup
		x -= lastMult
		i++
	}
	fmt.Println(x)
}
