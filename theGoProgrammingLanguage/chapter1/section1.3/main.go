package main

import (
	"bufio"
	"fmt"
	"os"
)

// exercise 1.3
// modify dup2 to print the names of all files in which each duplicated line occurs

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fprintf, err1 := fmt.Printf("dup2: %v\n", err)
				if err1 != nil {
					fmt.Printf("err1: %v\n", err1)
				}
				fmt.Println(fprintf)
				if err1 != nil {
					return
				}
				continue
			}
			countLines(f, counts)
			e := f.Close()
			if e != nil {
				return
			}
		}
	}
	for k, v := range counts {
		fmt.Printf("%v\n", k)
		for j, n := range v {
			if n > 1 {
				fmt.Printf("\t%v -> %v\n", n, j)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		fileName := f.Name()
		if counts[fileName] == nil {
			counts[fileName] = make(map[string]int)
		}
		counts[fileName][line]++
	}
}
