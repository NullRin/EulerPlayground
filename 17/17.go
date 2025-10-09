package main

import "fmt"

func main() {
	var totalSum = 0
	var ones = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var teens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	var tens = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	var i = 0
	var onesSums = 0
	for i < len(ones) {
		onesSums += len(ones[i])
		i++
	}
	i = 0
	var teensSums = 0
	for i < len(teens) {
		teensSums += len(teens[i])
		i++
	}
	i = 0
	var tensSums = 0
	for i < len(tens) {
		tensSums += len(tens[i])
		i++
	}
	i = 0
	var tensAndOnesSums = 0
	for i < len(tens) {
		tensAndOnesSums += (len(tens[i]) * 9) + onesSums
		i++
	}
	i = 0
	var hundredsSums = 0
	for i < len(ones) {
		// "{ones} hundred" = ones[i] + 7
		hundredsSums += len(ones[i]) + 7
		i++
	}
	i = 0
	var lessThan100Sums = onesSums + teensSums + tensSums + tensAndOnesSums
	var hundredsAndSums = 0
	for i < len(ones) {
		// "{ones} hundred and" = ones[i] + 10
		hundredsAndSums += len(ones[i]) + 10
		i++
	}
	var hundredsAndLessThan100Sums = hundredsAndSums + (9 * lessThan100Sums)
	var oneThousand = 11
	totalSum = hundredsSums + lessThan100Sums + hundredsAndLessThan100Sums + oneThousand
	fmt.Println(totalSum)

}
