package main

import (
	"bufio"
	"fmt"
	"learningGoAgain/theGoProgrammingLanguage/chapter2/ex2.1/tempconv"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Enter a number in fahrenheit and it will be converted to Celsius and Kelvin...")

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\r')
		if err != nil {
			fmt.Printf("ex2.2: %v", err)
			os.Exit(1)
		}
		d, err := strconv.ParseFloat(strings.TrimSpace(text), 32)
		if err != nil {
			fmt.Printf("ex2.2: %v", err)
			os.Exit(2)
		}
		f := tempconv.Fahrenheit(d)
		c := tempconv.FToC(f)
		k := tempconv.CToK(c)

		s := fmt.Sprintf("F: %.2f\t C: %.2f\t K:%.2f\t", f, c, k)
		fmt.Println(s)
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("ex2.2:%v", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.FToC(f)
		k := tempconv.CToK(c)

		s := fmt.Sprintf("F: %.2f\t C: %.2f\t K:%.2f\t", f, c, k)
		fmt.Println(s)
	}
}
