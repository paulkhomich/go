package main

func main() {
	println(join(" ", "my", "name", "is", "Pavel"))
	println(join("_", "variadic", "function", "test"))
}

func join(sep string, list ...string) string {
	var result string

	for i, v := range list {
		if i != 0 {
			result += sep
		}
		result += v
	}

	return result
}
