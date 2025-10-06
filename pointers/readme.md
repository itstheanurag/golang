# Go Pointers Example

This project demonstrates the basics of **pointers** in Go, including how to declare pointers, assign memory addresses, dereference pointers, and how pointers differ from passing values directly.

---

## What does the code do?

- Declares a pointer to an integer (`var ptr *int`) and prints its default value (`nil`).
- Creates an integer variable `myNumber` and stores its memory address in another variable (`pointer`).
- Shows how to access the value at the memory address using the asterisk (`*`) operator.
- Modifies the value at the memory address via the pointer, demonstrating that changes via the pointer affect the original variable.
- Prints the results to show the effect of pointer operations.

---

## Key Concepts

- `*` (asterisk) is used to declare a pointer type and to dereference a pointer (access the value at the address).
- `&` (ampersand) is used to get the memory address of a variable.

---

## Example Output

```
This is about the pointers
<nil>
pointer variable will hold the memory refrence of myNumber variable  0xc0000140a8
to see the value of myNumber variable we have to use astric  23
after operation value of pointer variable  26
value of myNumber after operations  26
```

---

## Passing by Value vs Passing by Reference

### Passing by Value (default in Go)

When you pass a variable to a function, Go passes a **copy** of the value:

```go
func incrementValue(val int) {
    val = val + 1
}

func main() {
    num := 10
    incrementValue(num)
    fmt.Println(num) // Output: 10 (unchanged)
}
```

### Passing by Reference (using pointers)

When you pass a pointer, the function can modify the original value:

```go
func incrementPointer(val *int) {
    *val = *val + 1
}

func main() {
    num := 10
    incrementPointer(&num)
    fmt.Println(num) // Output: 11 (changed)
}
```

---

## Summary

- **Pointers** allow you to reference and modify the original value stored in memory.
- Use `*` to declare and dereference pointers, and `&` to get the address of a variable.
- Passing by reference (using pointers) enables functions to modify the original variable, while passing by value only modifies a copy.

---

## File

- [`pointers/main.go`](pointers/main.go)
