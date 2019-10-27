package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "garden"
	s2 := "danger"
	s3 := "garage"
	
	fmt.Printf("%s\t%s\t%t\n", s1, s2, isAnagram(s1, s2))
	fmt.Printf("%s\t%s\t%t\n", s1, s3, isAnagram(s1, s3))
}

func isAnagram(s1, s2 string) bool {
	s1x, s2x := strings.Split(s1, ""), strings.Split(s2, "")
	sort.Strings(s1x)
	sort.Strings(s2x)

	return strings.Join(s1x, "") == strings.Join(s2x, "")
}
