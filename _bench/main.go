package main

func main() {
	println(sum(4, 9))
}

func sum(a, b int) (res int) {
	defer func() { res += 100 }()

	return a + b
}
