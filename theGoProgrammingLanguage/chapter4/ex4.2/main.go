// Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag
// to print the SHA384 or SHA512 hash instead.

package main

import (
	"bufio"
	"crypto/sha256"
	sha5122 "crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

type hashType int

const (
	SHA256 hashType = iota
	SHA384
	SHA512
)

var (
	sha384 = flag.Bool("sha384", false, "enable if want to use SHA384")
	sha512 = flag.Bool("sha512", false, "enable if want to use SHA512")
)

func main() {
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string...")
	string, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var chosenHash int
	if *sha384 == true {
		chosenHash = 1
	} else if *sha512 == true {
		chosenHash = 2
	} else {
		chosenHash = 0
	}

	switch chosenHash {
	case 0:
		fmt.Println("Retrieving sum using SHA256")
		hashedString := sha256.Sum256([]byte(string))
		fmt.Println(hashedString)
		fmt.Printf("\n%T", hashedString)
	case 1:
		fmt.Println("Retrieving sum using SHA384")
		hashedString := sha5122.Sum384([]byte(string))
		fmt.Println(hashedString)
		fmt.Printf("\n%T", hashedString)

	case 2:
		fmt.Println("Retrieving sum using SHA512")
		hashedString := sha5122.Sum512([]byte(string))
		fmt.Println(hashedString)
		fmt.Printf("\n%T", hashedString)
	}
}
