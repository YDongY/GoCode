package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	/**
	复数 i = 根号 -1

	则 i^2 = -1 i^3 = -i i^4 = 1...

	*/

	c := 3 + 4i

	fmt.Println(cmplx.Abs(c)) // 5

	// 欧拉公式 : e^(iπ) + 1 = 0
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // (0+1.2246467991473515e-16i)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         // (0+1.2246467991473515e-16i)
	fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1) // (0.000+0.000i)
}
