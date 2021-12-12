package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkFile(f string) {
	if _, err := os.Stat(f); err == nil {
		log.Println("File exists. Continuing...")
	} else {
		fmt.Println("File doesn't seem to exist.")
		log.Panic(err)
	}
}

func countLines(f string) {
	checkFile(f)

	file, err := os.Open(f)
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(file)

	isMultiLine := false
	totalLines := 0
	numComments := 0
	for scanner.Scan() {
		totalLines += 1
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		lastWord := splitLine[len(splitLine)-1]
		firstTwoChars := splitLine[0][:2]
		lastTwoChars := lastWord[len(lastWord)-2:]

		if firstTwoChars == "//" && isMultiLine == false {
			//fmt.Println("+1 comment")
			numComments += 1
		} else if firstTwoChars == "/*" {
			//fmt.Println("started multiline")
			isMultiLine = true
			numComments += 1
		} else if lastTwoChars == "*/" {
			//fmt.Println("ended multiline")
			isMultiLine = false
			numComments += 1
		} else if isMultiLine == true {
			numComments += 1
		}
	}
	fmt.Printf("Total lines found is: %v.\nnumComments is: %v.\n", totalLines, numComments)
	fmt.Printf("Lines of valid code found: %v", totalLines-numComments)

}

func main() {
	file := "D:\\Documents\\go_workspace\\src\\learningGoAgain\\countLines\\file1.txt"

	countLines(file)
}
