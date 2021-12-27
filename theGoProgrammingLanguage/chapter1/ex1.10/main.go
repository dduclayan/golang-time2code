package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func main() {
	outFile, err := os.Create("D:\\Documents\\go_workspace\\src\\learningGoAgain\\theGoProgrammingLanguage\\chapter1\\ex1.10\\log.txt")
	check(err)

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		data := <-ch
		// fmt.Println(<-ch) // receive from channel ch
		w, err := outFile.WriteString(data + "\n")
		check(err)
		fmt.Printf("Wrote %d bytes\n", w)
	}
	fmt.Fprintf(outFile, "%.2fs elapsed", time.Since(start).Seconds())
	outFile.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
