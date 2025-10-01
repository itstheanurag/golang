package main

import "fmt"

func main() {

	fmt.Println("This is about slices in Golang")

	var fruits = []string{"Apple", "Banana", "Mango"}
	fmt.Printf("The type of fruits is %T \n", fruits)
	fmt.Println("the fruits list is", fruits)
	fmt.Println("the length of fruits list is", len(fruits))

	fruits = append(fruits, "Kiwi", "Peach", "Strawberry")
	fmt.Println(fruits)
	fmt.Println("the length of fruits list is", len(fruits))

	// slicing the slice, 1 is the starting point, 4 will not be inclued in the result
	fruits = append(fruits[1:4], "New")
	fmt.Println(fruits)

	

}

/*

   Slices under the hood are just arrays. But they are much more powerful and majority of time used in Golang.

   If you are able to understand slices, you will be able to play with databases more easily.

   The initialization of a slice is same as of an array, with a difference that in slices you don't have to pass the number of elements that are going to come in it. But you have to put in the value at the time of declarations
*/
