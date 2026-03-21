package main

import (
	"fmt"
	"nandcep/go-gawf-example/math"
)

func main() {
	calc := math.Calculation{
		NumA: 8,
		NumB: 4,
	}
	fmt.Printf("Add Result = %f\n", calc.Add())
	fmt.Printf("Min Result = %f\n", calc.Min())
	fmt.Printf("Mul Result = %f\n", calc.Mul())
	fmt.Printf("Div Result = %f\n", calc.Div())
}
