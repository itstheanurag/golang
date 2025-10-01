package main

import "fmt"

const LogginToken string = "Variables declared with capital letters are Equivalent to be public \n and can be used anywhere in this program."

// unused variables gives errors
func main() {

	// string can only be with Double Quotes
	var username string = "Gaurav"
	fmt.Println(username)
	fmt.Printf("The type of the username variable is: %T \n", username)

	var unassigned string
	fmt.Println("this is default unassigned string value: ", unassigned)
	fmt.Printf("The type of the username variable is: %T \n", unassigned)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of the username variable is: %T \n", isLoggedIn)

	var smallInt uint8 = 2
	fmt.Println(smallInt)
	fmt.Printf("The type of the username variable is: %T \n", smallInt)

	// its always 0
	var unassignedNumber int
	fmt.Println(unassignedNumber)
	fmt.Printf("The type of the username variable is: %T \n", unassignedNumber)

	fmt.Println(LogginToken)
	fmt.Printf("The type of the username variable is: %T \n", LogginToken)

	// implicit type

	var implicitVariable = "This is going to be a string"
	fmt.Println(implicitVariable)
	fmt.Printf("The type of the username variable is: %T \n", implicitVariable)

	// implicitVariable = 2 is not allowed

	// no var style, variables can be declared without keyword var inside the methods only

	something := "This is fine"
	fmt.Println(something)
	fmt.Printf("The type of the username variable is: %T \n", something)

	// the type of this type of variable is int32 and its value is 97
	variable := 'a'
	fmt.Println(variable)
	fmt.Printf("The type of the username variable is: %T \n", variable)

}
