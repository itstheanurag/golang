package main

import "fmt"

func main () {
	appendToMainText := appendText()
	appendToMainText("Hello there you son of a bihhhh")
	appendToMainText("\nHope you are doing well all is good here")
	fmt.Println(appendToMainText("\nYes i have been doing great, you son of a bihhhh"))
}


func appendText() func(string) string{
	txt := "";

	return func(word string) string {
		txt += word
		return txt;
	}
}