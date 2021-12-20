/*
FileB that contains usernames and their status
Username IsActive
dwsauder Yes
birddog No
jjlee Yes
oevans Yes
chrisj No
mccurley No
lchin Yes

if active users are not found in FileA, you use your company main line number (444) 123-1233

input: FileA, FileB
output: active users and their phone numbers in format
username (xxx) xxx-xxxx

plan:
- read thru fileB, create a slice of users who have 'IsActive' set to yes
- read thru fileA, split line by first space and assign first split to var email and second split to var phoneNum
- if user from activeUser slice is found in the line, print username and ph
- if user from activeUser slice not found in the line, print username and default ph
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type file struct {
	filepath string
}

func getActiveUsers(fb file) []string {
	var activeUserList []string

	f, err := os.Open(fb.filepath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	scanner.Scan() // skips header row aka [Username IsActive]
	for scanner.Scan() {
		line := scanner.Text()
		splitLines := strings.Fields(line)
		user := splitLines[0]
		status := splitLines[1]

		if status == "Yes" {
			activeUserList = append(activeUserList, user)
		}
	}
	err = f.Close()
	if err != nil {
		return nil
	}

	return activeUserList

}

func printActiveUsers(ul []string) {
	fileA := file{filepath: "D:\\Documents\\go_workspace\\src\\learningGoAgain\\activeUsers\\fileA.log"}
	var foundUserInList []string
	defaultPh := "(444) 123-1233"

	f, err := os.Open(fileA.filepath)
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.SplitN(line, " ", 2)
		fullEmail := splitLine[0]
		ph := splitLine[1]
		user := strings.Split(fullEmail, "@")[0]

		exists := nameInList(user, ul)
		if exists == true {
			//fmt.Println("ok is true. value = ", exists)
			fmt.Printf("%v\t%v\n", user, ph)
			foundUserInList = append(foundUserInList, user)
		}
	}

	for _, val := range ul {
		exists := nameInList(val, foundUserInList)

		if exists == false {
			fmt.Printf("%v\t%v\n", val, defaultPh)
		}
	}

	err = f.Close()
	if err != nil {
		return
	}

}

func nameInList(n string, list []string) bool {
	exists := false
	for _, val := range list {
		if n == val {
			exists = true
		}
	}
	return exists
}

func main() {
	fileB := file{filepath: "D:\\Documents\\go_workspace\\src\\learningGoAgain\\activeUsers\\fileB.log"}
	activeUserList := getActiveUsers(fileB)
	printActiveUsers(activeUserList)
}
