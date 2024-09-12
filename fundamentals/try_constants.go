package main

import (
	"fmt"
	"math"
)
const s string = "constant" // const declares a constant value

func tryConstants() {
	fmt.Println(s)

	const n = 500000 // a const statement can appear anywhere a var can

	const d = 3e20 / n // const expressions performs arithmetic with arbitrary precision
	fmt.Println(d)

	fmt.Println(int64(d)) //a numeric constant has no type until its given one such as by an explicit conversion

	fmt.Println(math.Sin(n))
}