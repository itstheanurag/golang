# Go User Input Example

This project demonstrates how to read user input from the terminal in Go, process it, and convert it from a string to a number.

---

## What does the code do?

- Prints a welcome message.
- Prompts the user to provide a rating.
- Reads the user's input from the terminal as a string.
- Prints the input and its type.
- Converts the input string to a floating-point number.
- If the conversion is successful, adds 1 to the rating and prints the result.
- If the conversion fails, prints the error.

---

## Key Concepts

- **Reading Input:**  
  Uses `bufio.NewReader(os.Stdin)` and `ReadString('\n')` to read a line of input from the user.

- **String Trimming:**  
  Uses `strings.TrimSpace(input)` to remove any leading/trailing whitespace (including the newline character).

- **String to Number Conversion:**  
  Uses `strconv.ParseFloat` to convert the trimmed string to a `float64`.

- **Error Handling:**  
  Checks if the conversion was successful and handles errors appropriately.

---

## Example Run

```
This is a welcome message
Please provide a rating for us:
8.5
The type of the username variable is string
Added 1 to the rating  9.5
```

---

## File

- [`input/main.go`](main.go)
