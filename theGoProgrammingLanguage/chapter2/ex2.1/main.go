package main

import (
	"fmt"
	"learningGoAgain/theGoProgrammingLanguage/chapter2/ex2.1/tempconv"
)

var c tempconv.Celsius
var c1 tempconv.Celsius

func main() {
	c = 10.0
	c1 = 13.0

	x := tempconv.CToF(c)
	y := tempconv.CToK(c1)

	fmt.Println(x)
	fmt.Println(y)
}
