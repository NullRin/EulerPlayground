package main

import "fmt"

func main() {
	var start uint32 = 1000000
	var highestValue uint32 = 0
	var highestCount uint32 = 0
	var tempStart uint32 = 0
	var tempCount uint32 = 0

	for start > 0 {
		tempStart = start
		tempCount = 0
		for tempStart != 1 {
			tempCount++
			// odd if true
			if (tempStart & 1) == 1 {
				tempStart = 3*tempStart + 1
			} else {
				tempStart = tempStart >> 1
			}
		}
		if tempCount > highestCount {
			highestValue = start
			highestCount = tempCount
		}
		start--
	}
	fmt.Println(highestValue)
}
