package main

import "fmt"

func tryVarsFunction() {
	var a = "initial"
	fmt.Println(a)

	var b, c int =  1, 2 // declaring multiple variables at once
	fmt.Println(b, c)

	var d = true // boolean
	fmt.Println(d)

	var e int // prints 0 by defualt
	fmt.Println(e)

	f := "apple" // syntax short-hand
	fmt.Println(f)
}