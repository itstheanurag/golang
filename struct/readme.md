# Go Interfaces and Structs: Shapes Example

This project demonstrates how **interfaces** and **structs** work together in Go using a geometric shapes example. It covers how to define interfaces, implement them with structs, use type assertions, and leverage polymorphism.

---

## Key Concepts

### Structs

A **struct** in Go is a composite type that groups together fields (variables) under a single type. Structs are used to model real-world entities with properties.

**Example:**

```go
type Rectangle struct {
    Width  float64
    Height float64
}
```

### Interfaces

An **interface** defines a set of method signatures (behavior). Any type that implements all the methods of an interface is said to satisfy that interface, implicitly.

**Example:**

```go
type Shape interface {
    Area() float64
    Perimeter() float64
    Name() string
}
```

---

## How the Code Works

1. **Shape Interface:**  
   The `Shape` interface requires three methods: `Area()`, `Perimeter()`, and `Name()`.

2. **Struct Implementations:**

   - `Rectangle`, `Circle`, and `Triangle` structs each implement all methods of the `Shape` interface.
   - This means you can use any of these types wherever a `Shape` is expected.

3. **Polymorphism:**

   - You can create a slice of `Shape` and store different shape types in it.
   - Functions like `PrintShapeInfo` and `TotalArea` can operate on any `Shape`, regardless of the underlying type.

4. **Type Assertion and Type Switch:**
   - The code uses type assertion (`s.(Circle)`) and type switch to determine the concrete type of a `Shape` at runtime.

---

## Example Output

```
Rectangle Properties:
  Area: 50.00
  Perimeter: 30.00
  Dimensions: 10.00 x 5.00

Circle Properties:
  Area: 153.94
  Perimeter: 43.98
  Radius: 7.00

Triangle Properties:
  Area: 6.00
  Perimeter: 12.00
  Sides: 3.00, 4.00, 5.00

=== Collection of Shapes ===
1. Rectangle - Area: 50.00
2. Circle - Area: 153.94
3. Triangle - Area: 6.00

Total area of all shapes: 209.94

=== Type Checking ===
This is a Circle with radius: 7.00
It's a Circle with radius 7.00!
```

---

## Why Use Interfaces and Structs?

- **Structs** let you model data with fields.
- **Interfaces** let you define and enforce behavior.
- Together, they enable **polymorphism**: writing code that works with any type that satisfies an interface, making your programs more flexible and extensible.

---

## Summary Table

| Concept      | Purpose                                      | Example Syntax                  |
| ------------ | -------------------------------------------- | ------------------------------- |
| Struct       | Group related data fields                    | `type Rectangle struct { ... }` |
| Interface    | Define a set of method signatures (behavior) | `type Shape interface { ... }`  |
| Method       | Attach behavior to a struct                  | `func (r Rectangle) Area() {}`  |
| Polymorphism | Write code that works with many types        | `func PrintShapeInfo(s Shape)`  |

---

## File

- [`struct/main.go`](struct/main.go)
