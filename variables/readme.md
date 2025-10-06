# Go Variable Types and Their Use Cases

This project demonstrates the declaration and usage of different variable types in Go, including their syntax, default values, and typical use cases.

---

## Basic Types in Go

### 1. **String**

- **Declaration:** `var username string = "Gaurav"`
- **Default Value:** `""` (empty string)
- **Use Case:** Storing text data.
- **Note:** Strings must be enclosed in double quotes.

### 2. **Boolean**

- **Declaration:** `var isLoggedIn bool = false`
- **Default Value:** `false`
- **Use Case:** Representing true/false conditions, flags, or state.

### 3. **Integer Types**

- **Signed Integers:** `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned Integers:** `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Declaration:** `var smallInt uint8 = 2`
- **Default Value:** `0`
- **Use Case:** Counting, indexing, storing numeric data. Use unsigned types when negative values are not needed.

### 4. **Constants**

- **Declaration:** `const LogginToken string = "..."`
- **Use Case:** Values that should not change during program execution. Constants with capital letters are exported (public).

### 5. **Implicit Typing**

- **Declaration:** `var implicitVariable = "This is going to be a string"`
- **Use Case:** Letting Go infer the type based on the assigned value.

### 6. **Short Variable Declaration**

- **Declaration:** `something := "This is fine"`
- **Use Case:** Concise variable declaration inside functions.

### 7. **Rune (alias for int32)**

- **Declaration:** `variable := 'a'`
- **Value:** Stores Unicode code points (e.g., `'a'` is 97)
- **Use Case:** Representing characters.

---

## Example Output

```
Gaurav
The type of the username variable is: string
this is default unassigned string value:
The type of the username variable is: string
false
The type of the username variable is: bool
2
The type of the username variable is: uint8
0
The type of the username variable is: int
Variables declared with capital letters are Equivalent to be public
 and can be used anywhere in this program.
The type of the username variable is: string
This is going to be a string
The type of the username variable is: string
This is fine
The type of the username variable is: string
97
The type of the username variable is: int32
```

---

## Summary Table

| Type   | Example Declaration      | Default Value | Use Case                       |
| ------ | ------------------------ | ------------- | ------------------------------ |
| string | `var s string = "hello"` | `""`          | Text data                      |
| bool   | `var b bool = true`      | `false`       | Flags, conditions              |
| int    | `var i int = 42`         | `0`           | General integers               |
| uint8  | `var u uint8 = 255`      | `0`           | Small positive integers        |
| rune   | `var r rune = 'a'`       | `0`           | Unicode code points/characters |
| const  | `const Pi = 3.14`        | N/A           | Unchanging values              |
| :=     | `x := 10`                | N/A           | Short declaration in functions |

---

## File

- [`variables/main.go`](variables/main.go)
