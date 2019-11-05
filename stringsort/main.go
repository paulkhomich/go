package main

import "sort"

func main() {
	r := []rune("helloworld")
	sort.Sort(stringSorted(r))
	println(string(r))
}
