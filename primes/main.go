package main

import (
	"fmt"
	"math"
	"time"
)

func sieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	var primeNumbers []int
	for i, isPrime := range primes {
		if isPrime {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}

func main() {
	n := 10000000
	start := time.Now()
	primeNumbers := sieveOfEratosthenes(n)
	elapsed := time.Since(start)
	fmt.Printf("Primes found: %d, last prime: %d\n", len(primeNumbers), primeNumbers[len(primeNumbers)-1])
	fmt.Printf("Time taken: %s\n", elapsed)
}
