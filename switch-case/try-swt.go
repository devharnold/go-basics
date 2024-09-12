package main


import (
	"fmt"
	"time"
)


func main() {
	// here we are going to implement a basic switch in Go
	i := 2
	fmt.Println("Write ", i, "as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	// used commas to seperate the multiple expressions and the optional default case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend")
	default:
		fmt.Println("Its a weekday!")
	}

	// switch wothout expressions is another way to express if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	// a type switch compares types instead of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(5)
	whatAmI("hey")
}