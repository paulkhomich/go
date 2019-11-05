package main

type runeSorted []rune

func (s runeSorted) Len() int { return len(s) }
func (s runeSorted) Less(i, j int) bool { return s[i] < s[j] }
func (s runeSorted) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
