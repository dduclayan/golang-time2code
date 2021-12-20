package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

type ByCount []WordCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func topTenWords(filePath string) map[string]int {
	words := make(map[string]int)
	file, err := os.Open(filePath)
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)
		for _, v := range splitLine {
			if _, ok := words[v]; ok {
				words[v]++
			} else {
				words[v] = 1
			}
		}
	}
	return words
}

func main() {
	filePath := "C:\\Users\\Devin Duclayan\\AppData\\Roaming\\JetBrains\\GoLand2021.2\\scratches\\some_file.log"
	// flagFilePath := flag.String("filePath", "", "Path to file")
	// flag.Parse()

	topWords := topTenWords(filePath)

	p := make(ByCount, len(topWords))
	// fmt.Println(p)

	i := 0
	for k, v := range topWords {
		p[i] = WordCount{k, v}
		fmt.Println(p)
		i++
	}

	sort.Sort(p)

	for _, k := range p[:10] {
		fmt.Printf("%v\t\t%v\n", k.Word, k.Count)
	}
}
