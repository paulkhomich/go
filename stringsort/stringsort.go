package main

type stringSorted []rune

func (s stringSorted) Len() int { return len(s) }
func (s stringSorted) Less(i, j int) bool { return s[i] < s[j] }
func (s stringSorted) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
