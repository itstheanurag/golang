# Go Channels - Comprehensive Guide

## What Are Channels?

Channels are Go's built-in mechanism for **communication between goroutines**. They provide a way to safely pass data between concurrent operations without using locks or shared memory.

Think of a channel as a **pipe** through which you can send and receive values of a specific type.

## Syntax

```go
// Create an unbuffered channel
ch := make(chan int)

// Create a buffered channel with capacity 5
ch := make(chan int, 5)

// Send to channel
ch <- 42

// Receive from channel
value := <-ch

// Close a channel
close(ch)
```

## Types of Channels

### 1. **Unbuffered Channels**

```go
ch := make(chan int)
```

- No capacity to hold values
- Send operation blocks until receiver is ready
- Receive operation blocks until sender sends
- Guarantees **synchronization** between sender and receiver

### 2. **Buffered Channels**

```go
ch := make(chan int, 3)
```

- Has capacity to hold N values
- Send blocks only when buffer is full
- Receive blocks only when buffer is empty
- Allows **asynchronous** communication

### 3. **Directional Channels**

```go
// Send-only channel
func sender(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receiver(ch <-chan int) {
    value := <-ch
}
```

- Improves type safety
- Prevents misuse (sending on receive-only channel, etc.)

## Key Operations

### Sending

```go
ch <- value  // Blocks if channel is full (buffered) or has no receiver (unbuffered)
```

### Receiving

```go
value := <-ch        // Blocks until value is available
value, ok := <-ch    // ok is false if channel is closed and empty
```

### Closing

```go
close(ch)  // Only sender should close channels
```

**Important**: Closing a channel signals "no more values will be sent"

- Receiving from a closed channel returns zero value immediately
- Sending to a closed channel causes panic
- Closing a closed channel causes panic

### Range Over Channel

```go
for value := range ch {
    fmt.Println(value)
}
// Loop terminates when channel is closed
```

### Select Statement

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case ch3 <- 42:
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No channel ready")
}
```

## Common Patterns

### 1. **Worker Pool**

Multiple workers process jobs from a shared


## File

- [`channels/main.go`](main.go)