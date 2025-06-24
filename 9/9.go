package main

import "fmt"

func main() {
	// a=m^{2}-n^{2}, b=2mn, c=m^{2}+n^{2} ------ a^2+b^2=c^2
	var target = 1000
	var mStart = 1
	var m, n int
	var a, b, c int
	for 2*m < target {
		mStart++
		m = mStart
		n = mStart - 1
		a = 0
		b = 0
		c = 0
		for a+b+c < target {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			m++
		}
		if a+b+c == target {
			fmt.Println(a, b, c)
			break
		}
	}
	fmt.Println("End")
}
