/*
output: print out the time spent for each process (identified by pid)
000001 4 days, 4:30:45

plan:
- create a struct that stores startTime, endTime, runningTime
- add pid to map, have pid as key and struct as the value. [pid][]timeStruct{}
- store startTimes and endTimes as you go
- after adding all pids and times, calculate the running time and print
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	time2 "time"
)

type timeLog struct {
	startTime   time2.Time
	runningTime time2.Duration
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func clearThingsUp(f string) {
	file, err := os.Open(string(f))
	check(err)

	pidTime := make(map[string]timeLog)
	scanner := bufio.NewScanner(file)

	scanner.Scan() // skip header
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)

		time := splitLine[0]
		date := splitLine[1]
		pid := splitLine[2]
		//status := splitLine[3]

		layout := "15:04:05 1/02/2006"
		t, err := time2.Parse(layout, time+" "+date)
		if err != nil {
			log.Fatal(err)
		}

		if _, ok := pidTime[pid]; ok {
			pidTime[pid] = timeLog{
				startTime:   pidTime[pid].startTime,
				runningTime: t.Sub(pidTime[pid].startTime)}
		} else {
			pidTime[pid] = timeLog{startTime: t}
		}
	}
	sortAndPrint(pidTime)
}

func sortAndPrint(m map[string]timeLog) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, m[key].runningTime)
	}
}

func main() {
	fileA := "D:\\Documents\\go_workspace\\src\\learningGoAgain\\clearThingsUp\\log1.txt"

	clearThingsUp(fileA)
}
