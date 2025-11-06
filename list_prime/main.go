package main

import "fmt"

var isPrime [1_000_000]bool

func main() {
	for i := 2; i < len(isPrime); i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < len(isPrime); i++ {
		if isPrime[i] {
			for j := i * i; j < len(isPrime); j += i {
				isPrime[j] = false
			}
		}
	}

	for i, p := range isPrime {
		if p {
			fmt.Printf("%d\n", i)
		}
	}
}
