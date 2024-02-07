package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / 2 * z
		fmt.Printf("Iteration %d: z = %v\n", i+1, z)
	}
	return z
}

func main() {
	fmt.Println("Square root of number: %v", Sqrt(2))
}
