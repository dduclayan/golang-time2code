package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "anagram"
	s2 := "nagaram"

	sorted1 := sortString(s1)
	sorted2 := sortString(s2)

	fmt.Println(isAnagram(sorted1, sorted2))

}

func isAnagram(x, y string) bool {
	if x == y {
		return true
	}
	return false
}

func sortString(s string) string {
	w := strings.Split(s, "")
	sort.Strings(w)
	return strings.Join(w, "")

}
