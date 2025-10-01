package main

import "fmt"

func main() {
	fmt.Println("This is about the pointers")

	var ptr *int
	fmt.Println(ptr)

	// lets create a variable and store its memory address in another variable

	myNumber := 23

	var pointer = &myNumber
	fmt.Println("pointer variable will hold the memory refrence of myNumber variable ", pointer)
	fmt.Println("to see the value of myNumber variable we have to use astric ", *pointer)

	*pointer = *pointer + 3

	fmt.Println("after operation value of pointer variable ", *pointer)
	fmt.Println("value of myNumber after operations ", myNumber)
}

/*
* All programming languages have a problem,
  when you pass a variable to a function there is no gaurantee
  that actual value will be passed in the function not the copies of the those
  variables.

  Pointers solves this problem and make sure that everytime a memory address refrence is passed
  instead of copies of the said variables.

  * (astric) sign is used to make a variable as pointer
  & (ampresant) is refrences the memory address of a pointer
*/
