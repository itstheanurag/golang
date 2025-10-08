# Go Loops Comprehensive Guide

This project demonstrates **all major types of loops and loop-related features in Go**. Each example in `loops/main.go` is annotated and explained, making it a practical reference for Go developers.

---

## 1. Basic For Loop (C-style)

```go
for i := 0; i < 5; i++ { ... }
```

- Standard loop with initialization, condition, and post statement.
- Used for a known number of iterations.

---

## 2. While-Style Loop

```go
for condition { ... }
```

- Only a condition; acts like a `while` loop in other languages.
- Runs as long as the condition is true.

---

## 3. Infinite Loop with Break

```go
for { ... }
```

- Runs forever unless a `break` statement is encountered.

---

## 4. Range Over Slice

```go
for index, value := range slice { ... }
```

- Iterates over elements of a slice, providing both index and value.

---

## 5. Range Over Slice (Index Only)

```go
for index := range slice { ... }
```

- Iterates over indices only.

---

## 6. Range Over Slice (Value Only)

```go
for _, value := range slice { ... }
```

- Ignores the index, iterates over values only.

---

## 7. Range Over Map

```go
for key, value := range map { ... }
```

- Iterates over key-value pairs in a map.

---

## 8. Range Over String

```go
for index, char := range string { ... }
```

- Iterates over Unicode code points (runes) in a string.

---

## 9. Continue Statement

```go
for ... {
    if condition { continue }
    ...
}
```

- Skips the rest of the current iteration and moves to the next.

---

## 10. Break Statement

```go
for ... {
    if condition { break }
    ...
}
```

- Exits the loop immediately.

---

## 11. Nested Loops

```go
for ... {
    for ... {
        ...
    }
}
```

- Loops inside loops, e.g., for creating multiplication tables.

---

## 12. Labeled Break

```go
OuterLoop:
for ... {
    for ... {
        break OuterLoop
    }
}
```

- Breaks out of an outer loop from within an inner loop.

---

## 13. Labeled Continue

```go
OuterContinue:
for ... {
    for ... {
        continue OuterContinue
    }
}
```

- Skips to the next iteration of an outer loop.

---

## 14. Range Over Channel

```go
for value := range channel { ... }
```

- Iterates over values received from a channel until it is closed.

---

## 15. Range Over Array

```go
for i, num := range array { ... }
```

- Iterates over elements of an array.

---

## 16. Backwards Loop

```go
for i := n; i > 0; i-- { ... }
```

- Counts down instead of up.

---

## 17. Loop with Step of N

```go
for i := 0; i < n; i += step { ... }
```

- Increments by a value other than 1.

---

## 18. Loop with Multiple Variables

```go
for i, j := 0, 10; i < 5; i, j = i+1, j-1 { ... }
```

- Manages multiple loop variables.

---

## 19. Range Value is a Copy

- The value in a `range` loop is a copy; modifying it does not affect the original slice.

---

## 20. Modifying Slice Using Index

- Use the index to modify the original slice.

---

## 21. Pointer Issue in Range Loop

- If you take the address of the loop variable, all pointers may point to the same memory.
- Correct way: create a new variable inside the loop.

---

## 22. Early Return from Loop

- Use `return` to exit a function early from within a loop.

---

## 23. Goto Statement

```go
goto Label
```

- Jumps to a labeled statement. Rarely used, but available.

---

## Summary Table

| Loop Type               | Syntax Example                    | Use Case             |
| ----------------------- | --------------------------------- | -------------------- |
| Basic For               | `for i := 0; i < n; i++ {}`       | Fixed iterations     |
| While-Style             | `for condition {}`                | Unknown iterations   |
| Infinite                | `for {}`                          | Loop until break     |
| Range (slice/map/array) | `for k, v := range collection {}` | Collections          |
| Nested                  | `for ... { for ... {}}`           | Tables, grids        |
| Labeled Break/Continue  | `break Label` / `continue Label`  | Control nested loops |
| Channel                 | `for v := range ch {}`            | Concurrency          |
| Goto                    | `goto Label`                      | Rare, special cases  |

---

## File

- [`loops/main.go`](main.go)
