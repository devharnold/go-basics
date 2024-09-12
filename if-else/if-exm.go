package main

import "fmt"

// you'll notice in Go we dont have to use parentheses but curly braces must be there
// if statement in Go is straight-forward
func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// you can have the if statement without the else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//logical operators like && and || are often used in conditions
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 are even")
	}

	/*
	a statement can precede conditionals
	any variables declared in this statement are available in the current and subsequent branches
	*/
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}