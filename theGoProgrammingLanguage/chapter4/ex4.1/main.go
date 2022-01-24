// Write a func that counts the number of bits that are different in two SHA256 hashes

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	x := sha256.Sum256([]byte("x"))
	y := sha256.Sum256([]byte("X"))

	fmt.Println(x)
	fmt.Println(y)

	z := countDiffBits(x, y)
	fmt.Println(z)
}

func countDiffBits(x, y [32]byte) (diffBytes int) {
	diffBytes = 0
	for i, v := range x {
		if v != y[i] {
			fmt.Printf("%v != %v\n", v, y[i])
			diffBytes++
		} else {
			fmt.Printf("Found a match! %v == %v at index %v\n", v, y[i], i)
		}
	}
	return diffBytes
}
