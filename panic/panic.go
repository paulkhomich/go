package main

func main() {
	a := f()
	println(a)
}

func f() (ans int) {
	defer func () { 
		recover()
		ans = 1
	}()
	panic(0)
}