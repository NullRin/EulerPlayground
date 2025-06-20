package main

import (
	"fmt"
	"math"
	"sort"
)

type Palindrome struct {
	palindrome int32
	n1         int16
	n2         int16
}

func reverseNum(num int32) int32 {
	var reversedNum int32 = 0
	var remainder int32 = 0
	var cycles int8 = 1
	// fmt.Println(num)
	var digits = int8(math.Log10(float64(num)) + 1)
	for num > 0 {
		remainder = num % 10
		reversedNum += remainder * int32(math.Pow10(int(digits-cycles)))
		num = num / 10
		cycles++
	}
	// fmt.Println(reversedNum)
	return reversedNum
}

func main() {
	var n1 int16 = 999
	var n2 int16 = 999
	var n3 int32 = 0
	var palindromes []Palindrome
	for n2 >= 100 {
		n1 = 999
		for n1 >= 100 {
			n3 = int32(n1) * int32(n2)
			if n3 == reverseNum(n3) {
				palindromes = append(palindromes, Palindrome{n3, n1, n2})
			}
			n1--
		}
		n2--
	}
	sort.Slice(palindromes, func(i, j int) bool { return palindromes[i].palindrome < palindromes[j].palindrome })
	if len(palindromes) >= 1 {
		fmt.Println(palindromes)
	} else {
		fmt.Println("No palidroms found")
	}

}
