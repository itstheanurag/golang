package main

import (
	"fmt"
	"math"
	"time"
)

func sieveOfEratosthenes(n int) (int, int) {
	primes := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		primes[i] = true
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	var primesFound = 0
	var lastPrime = 2
	for i, isPrime := range primes {
		if isPrime {
			primesFound++
			lastPrime = i
		}
	}

	return primesFound, lastPrime
}

func main() {
	n := 100000000
	start := time.Now()
	primesFound, lastPrime := sieveOfEratosthenes(n)
	elapsed := time.Since(start)
	fmt.Printf("Primes found: %d, last prime: %d\n", primesFound, lastPrime)
	fmt.Printf("Time taken: %s\n", elapsed)
}
