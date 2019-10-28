package main

import "crypto/sha256"

func main() {
	a1 := sha256.Sum256([]byte("A"))
	a2 := sha256.Sum256([]byte("B"))

	println(shaDiff(&a1, &a2))
}

func shaDiff(a1, a2 *[32]byte) uint8 {
	var counter uint8 = 0
	for i := 0; i < 32; i++ {
		shaXor := a1[i] ^ a2[i]
		for j := 0; j < 8; j++ {
			counter += (shaXor >> j) & 1
		}
	}

	return counter
}
