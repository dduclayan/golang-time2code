package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//func main() {
//	var s, sep string
//	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(s)
//}

// exercise 1.1
// modify echo program to print os.Args[0]
func function1() {
	var s, sep string
	fmt.Println("Exercise 1.1.")
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("os.Args[0]: %v\n\n", os.Args[0])
}

// exercise 1.2
// Modify the echo program to print the index and value of each of its arguments, one per line
func function2() {
	fmt.Println("exercise 1.2")
	for i, v := range os.Args {
		fmt.Printf("index: %v\t value: %v\n", i, v)
	}
	fmt.Println()

}

// exercise 1.3
func function3() {
	start := time.Now()

	var s, sep string
	fmt.Println("Exercise 1.3")
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	end := time.Since(start).Microseconds()
	fmt.Printf("The above took %v microseconds.\n\n", end)

}

func function4() {
	start := time.Now()

	fmt.Println("Exercise 1.4")
	fmt.Println(strings.Join(os.Args[1:], " "))

	end := time.Since(start).Microseconds()
	fmt.Printf("The above took %v microseconds.\n\n", end)

}

func main() {
	function1()
	function2()
	function3()
	function4()
}
