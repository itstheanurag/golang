package main

import "fmt"

// main is a reserved keyword you can't give it to anything
// main func is the only function that is executed
func main() {
	fmt.Println("Hello World from the Go lang")
	output()
}

func output() {
	fmt.Println("Output func is executed")
}
