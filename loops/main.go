package main

import "fmt"

func main() {
	fmt.Println("=== GO LOOPS COMPREHENSIVE GUIDE ===\n")

	// 1. BASIC FOR LOOP (C-style)
	fmt.Println("1. Basic For Loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  Iteration %d\n", i)
	}

	// 2. WHILE-STYLE LOOP (condition only)
	fmt.Println("\n2. While-Style Loop:")
	count := 0
	for count < 3 {
		fmt.Printf("  Count is %d\n", count)
		count++
	}

	// 3. INFINITE LOOP with break
	fmt.Println("\n3. Infinite Loop with Break:")
	n := 0
	for {
		if n >= 3 {
			break
		}
		fmt.Printf("  n = %d\n", n)
		n++
	}

	// 4. RANGE OVER SLICE
	fmt.Println("\n4. Range Over Slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("  Index %d: %s\n", index, fruit)
	}

	// 5. RANGE OVER SLICE (index only)
	fmt.Println("\n5. Range Over Slice (index only):")
	for index := range fruits {
		fmt.Printf("  Index: %d\n", index)
	}

	// 6. RANGE OVER SLICE (value only, using blank identifier)
	fmt.Println("\n6. Range Over Slice (value only):")
	for _, fruit := range fruits {
		fmt.Printf("  Fruit: %s\n", fruit)
	}

	// 7. RANGE OVER MAP
	fmt.Println("\n7. Range Over Map:")
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	for name, age := range ages {
		fmt.Printf("  %s is %d years old\n", name, age)
	}

	// 8. RANGE OVER STRING (iterates over runes/characters)
	fmt.Println("\n8. Range Over String:")
	text := "Hello"
	for index, char := range text {
		fmt.Printf("  Index %d: %c (Unicode: %d)\n", index, char, char)
	}

	// 9. CONTINUE STATEMENT
	fmt.Println("\n9. Continue Statement (skip even numbers):")
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("  Odd number: %d\n", i)
	}

	// 10. BREAK STATEMENT
	fmt.Println("\n10. Break Statement (stop at 5):")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("  Number: %d\n", i)
	}

	// 11. NESTED LOOPS
	fmt.Println("\n11. Nested Loops (multiplication table):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("  %d x %d = %d\n", i, j, i*j)
		}
	}

	// 12. LABELED BREAK (break outer loop)
	fmt.Println("\n12. Labeled Break:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("  Breaking outer loop at i=1, j=1")
				break OuterLoop
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}

	// 13. LABELED CONTINUE (continue outer loop)
	fmt.Println("\n13. Labeled Continue:")
OuterContinue:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Printf("  Skipping rest of inner loop at i=%d, j=%d\n", i, j)
				continue OuterContinue
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}

	// 14. RANGE OVER CHANNEL
	fmt.Println("\n14. Range Over Channel:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // Must close channel for range to terminate
	for value := range ch {
		fmt.Printf("  Received: %d\n", value)
	}

	// 15. RANGE WITH ARRAY
	fmt.Println("\n15. Range Over Array:")
	numbers := [4]int{10, 20, 30, 40}
	for i, num := range numbers {
		fmt.Printf("  numbers[%d] = %d\n", i, num)
	}

	// 16. BACKWARDS LOOP
	fmt.Println("\n16. Backwards Loop:")
	for i := 5; i > 0; i-- {
		fmt.Printf("  Countdown: %d\n", i)
	}

	// 17. STEP BY N
	fmt.Println("\n17. Loop with Step of 2:")
	for i := 0; i < 10; i += 2 {
		fmt.Printf("  Even number: %d\n", i)
	}

	// 18. LOOP WITH MULTIPLE VARIABLES
	fmt.Println("\n18. Loop with Multiple Variables:")
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("  i=%d, j=%d\n", i, j)
	}

	// 19. RANGE VALUE IS A COPY
	fmt.Println("\n19. Range Value is a Copy (modifying doesn't change original):")
	nums := []int{1, 2, 3}
	for _, num := range nums {
		num = num * 2 // This doesn't modify the original slice
		fmt.Printf("  Modified copy: %d\n", num)
	}
	fmt.Printf("  Original slice: %v\n", nums)

	// 20. MODIFYING SLICE USING INDEX
	fmt.Println("\n20. Modifying Slice Using Index:")
	for i := range nums {
		nums[i] = nums[i] * 2 // This modifies the original slice
	}
	fmt.Printf("  Modified slice: %v\n", nums)

	// 21. POINTER IN RANGE LOOP
	fmt.Println("\n21. Pointer Issue in Range Loop:")
	items := []string{"a", "b", "c"}
	var pointers []*string
	
	// WRONG WAY - all pointers point to same variable
	for _, item := range items {
		pointers = append(pointers, &item) // Don't do this!
	}
	fmt.Println("  Wrong way (all point to last value):")
	for _, ptr := range pointers {
		fmt.Printf("    %s\n", *ptr)
	}
	
	// CORRECT WAY - create new variable
	pointers = nil
	for _, item := range items {
		item := item // Create new variable
		pointers = append(pointers, &item)
	}
	fmt.Println("  Correct way:")
	for _, ptr := range pointers {
		fmt.Printf("    %s\n", *ptr)
	}

	// 22. EARLY RETURN FROM LOOP
	fmt.Println("\n22. Early Return from Function:")
	result := findFirstEven([]int{1, 3, 5, 8, 9})
	fmt.Printf("  First even number: %d\n", result)

	// 23. GOTO (rarely used, but available)
	fmt.Println("\n23. Goto Statement:")
	i := 0
Loop:
	if i < 3 {
		fmt.Printf("  i = %d\n", i)
		i++
		goto Loop
	}
}

// Helper function for example 22
func findFirstEven(numbers []int) int {
	for _, num := range numbers {
		if num%2 == 0 {
			return num // Early return from loop
		}
	}
	return -1 // Not found
}