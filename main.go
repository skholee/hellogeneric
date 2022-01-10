package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	predicateOfInt := func(i int) bool { return i%2 == 0 }

	float64s := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	predicateOfFloat64 := func(f float64) bool { return f >= 5.0 }

	fmt.Printf("No-Generic filters: %v and %v\n",
		FilterInts(ints, predicateOfInt),
		FilterFloat64s(float64s, predicateOfFloat64))
}
