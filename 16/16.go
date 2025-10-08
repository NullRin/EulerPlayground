package main

import "fmt"

func main() {
	// manual multiplication
	// 2^1000 has about 300 digits
	var length uint16 = 350
	var sum = make([]uint8, length)
	var i uint16 = 1
	var j uint16 = 1
	var temp uint8
	var carry uint8
	var totalSum uint32 = 0
	sum[0] = 2
	for i < 1000 {
		j = 0
		temp = 0
		carry = 0
		for j < length-1 {
			temp = (sum[j] << 1) + carry
			carry = 0
			if temp >= 10 {
				temp = temp - 10
				carry = 1
			}
			sum[j] = temp
			j++
		}
		i++
	}
	i = 0
	for i < length {
		totalSum += uint32(sum[i])
		i++

	}
	fmt.Println(sum)
	fmt.Println(totalSum)
}
