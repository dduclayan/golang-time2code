package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	myTime := "9/25/2020 15:04"
	t, err := time.Parse("1/02/2006 15:04", myTime)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(t)

}
