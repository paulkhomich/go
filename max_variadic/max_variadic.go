package main

func main() {
	println(max(1))
	println(max(1,2,3,40,5,6))
	println(max(-100, 10, 100))
}

func max(val int, vals ...int) int {
	max := val
	for _, v := range vals {
		if v > max {
			max = v
		}
	}

	return max
}
