package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popCount returns the population count (number of set bits aka the '1's in a binary number like 101) of x
func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// my function utilizes a for loop instead
func myPopCount(x uint64) int {
	var y byte
	for i := 0; i < 8; i++ {
		y += pc[byte(x>>(i*8))]
	}
	return int(y)
}

func main() {
	x := popCount(12312312331343447)
	y := myPopCount(12312312331343447)
	fmt.Println(x)
	fmt.Println("mine: ", y)
}
