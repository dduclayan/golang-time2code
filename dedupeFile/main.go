/*
Given a file input with multiple lines and some duplicates.

Output a new file with duplicates removed.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func dedupeFile(filePath string) []string {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	lineCount := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if _, ok := lineCount[line]; ok {
		} else {
			lineCount[line] = true
		}
	}

	var dedupedStrings []string
	for k := range lineCount {
		dedupedStrings = append(dedupedStrings, k)
	}
	sort.Strings(dedupedStrings)
	fmt.Println(dedupedStrings)

	return dedupedStrings
}

func writeStrings(s []string) {
	outfile := "D:\\Documents\\go_workspace\\src\\learningGoAgain\\dedupeFile\\file2.txt"
	file, err := os.Create(outfile)
	check(err)
	defer file.Close()

	fmt.Println("Logging to file...")
	for _, v := range s {
		_, err := file.Write([]byte(v + "\n"))
		if err != nil {
			return
		}
	}
}

func main() {
	file := "D:\\Documents\\go_workspace\\src\\learningGoAgain\\dedupeFile\\logFile.txt"
	x := dedupeFile(file)
	writeStrings(x)
}
