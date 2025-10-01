package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "This is a welcome message"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	// fmt.Println(reader)
	fmt.Println("Please provide a rating for us: ")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("The type of the username variable is %T \n", input)

	// Now we have gotten a string from user keyboard.
	// but what if we want to convert that string to number
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating ", numRating+1)
	}

}
