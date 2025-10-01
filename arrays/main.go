package main

import "fmt"

func main() {
	fmt.Println("Arrays in the Golang")

	// An array of size 5 is initialized with type string
	var fruits [5]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[3] = "Banana"

	fmt.Println("the fruits list is", fruits)
	fmt.Println("the length of fruits list is", len(fruits))

	// this is how arrays are given values during declaration
	var vegetables = [4]string{"Potato", "Peas", "Carrot", "Cauliflower"}

	fmt.Println("the vegetables list is", vegetables)
	fmt.Println("the length of vegetables list is", len(vegetables))
}

/*
  Arrays are not much used in the Golang, they are the least used datastructure compared to other langauges.
  But behind the scenes go does use Arrays.

  we have to explicitly tell how much data is going to come inside the array. (its compulsory)

  There is not much in the arrays, there is no sorting, no filtering and other operations.

*/
