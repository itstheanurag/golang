# Sieve of Eratosthenes in Go

This project demonstrates the use of the **Sieve of Eratosthenes** algorithm to efficiently find all prime numbers up to a given number `n` in Go.

## How does the code work?

- The main logic is implemented in the `sieveOfEratosthenes` function.
- It creates a boolean slice to track which numbers are prime.
- It marks all numbers as prime initially, then iteratively marks multiples of each prime as not prime.
- After sieving, it collects all numbers still marked as prime into a slice and returns them.

## Main Steps

1. **Initialization:**  
   Create a boolean array `primes` of size `n+1`, set all entries from 2 to `n` as `true`.

2. **Sieving:**  
   For each number `i` from 2 up to the square root of `n`, if `i` is marked as prime, mark all multiples of `i` as not prime.

3. **Collecting Primes:**  
   Iterate through the boolean array and collect all indices marked as `true` (these are the primes).

4. **Timing and Output:**  
   In `main`, the code measures how long it takes to find all primes up to 10,000,000, prints the number of primes found, the last prime, and the time taken.

## Example Output

```
Primes found: 664579, last prime: 9999991
Time taken: 1.234567s
```

## File

- [`primes/main.go`](primes/main.go)

## References

- [Sieve of Eratosthenes - Wikipedia](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
