package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func httpCheck(s *string) {
	if strings.HasPrefix(*s, "https://") == false {
		*s = "https://" + *s
		fmt.Fprintf(os.Stdout, "'https://' was missing, added it to url on your behalf. Now url is: %v\n\n", *s)
	}
}

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Did not pass in a url. exiting...")
		os.Exit(0)
	}

	for _, url := range os.Args[1:] {
		httpCheck(&url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		resp.Body.Close()
	}
}
