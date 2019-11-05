package main

import "sort"

func main() {
	r := []rune("helloworld")
	sort.Sort(runeSorted(r))
	println(string(r))
}
