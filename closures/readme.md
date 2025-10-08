# Closures Example in Go

This example demonstrates the concept of **closures** in Go.

## What does the code do?

- The program defines a function `appendText` that returns another function (a closure).
- The closure captures and modifies a variable (`txt`) from its enclosing scope.
- In `main`, the closure is used to build up a string by appending new text each time it is called.

## How does it work?

1. `appendText` initializes an empty string variable `txt`.
2. It returns an **anonymous function** that takes a string argument, appends it to `txt`, and returns the updated string.
3. In `main`, `appendToMainText` is assigned the closure returned by `appendText`.
4. Each call to `appendToMainText` adds more text to `txt`, preserving its state between calls.
5. The final call prints the accumulated string.

## Example Output

```
Hello there you son of a bihhhh
Hope you are doing well all is good here
Yes i have been doing great, you son of a bihhhh
```

## What is a Closure?

A **closure** is a function that references variables from outside its own body. In this example, the anonymous function returned by `appendText` references and modifies the `txt` variable defined in `appendText`.

## File

- [`closures/main.go`](main.go)
