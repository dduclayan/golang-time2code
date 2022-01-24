package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "12312424234"
	c := comma(s)
	fmt.Println(c)
}

func comma(s string) string {
	var buf bytes.Buffer
	pre := len(s) % 3 // see if string is a multiple of 3
	if pre == 0 {
		pre = 3
	}

	buf.WriteString(s[:pre]) // assuming s = 12312424234. this would be '12'
	for i := pre; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3]) // write next three chars to buffer
	}
	return buf.String()
}
