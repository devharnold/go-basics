package main

import "fmt"

func tryForsFunc() {
	i := 3 // most basic type with a single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { // a classic/initial/condition after for loop
		fmt.Println(j)
	}
	
	for i := range 3 { // another way of accomplishing basic "do this N times" iteration is range over an integer
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop") // without a condition loops repeatedly until you break or return from the enclosing function
		break
	}

	// continue to the next iteration of the loop
	for n := range 6 {
		if n %2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}