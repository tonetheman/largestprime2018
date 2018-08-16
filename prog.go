package main

import (
	"fmt"
	"math/big"
	"time"
)

// from here
// https://academyofmathematics.wordpress.com/2018/08/15/the-largest-known-prime-number/

func main() {
	var a big.Int
	base := int64(2)
	exponent := int64(77232917)
	zero := int64(0)

	beforeExp := time.Now()
	a.Exp(big.NewInt(base), big.NewInt(exponent), big.NewInt(zero))
	eTime := time.Since(beforeExp)
	fmt.Println("elapsed time to computer", eTime)
	beforeStr := time.Now()
	s := a.String()
	sTime := time.Since(beforeStr)
	fmt.Println("elapsed time to create string", sTime)
	beforePr := time.Now()
	fmt.Println(s)
	pTime := time.Since(beforePr)
	fmt.Println("elapsed time to print", pTime)

}
