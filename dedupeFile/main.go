package main

import (
	"bufio"
	"log"
	"os"
)

func dedupeFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)
}

func main() {
	file := "D:\\Documents\\go_workspace\\src\\learningGoAgain\\dedupe_file\\logFile.txt"
	dedupeFile(file)
}
