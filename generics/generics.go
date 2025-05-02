package main

import "fmt"

func sumInts(l map[string]int64) int64 {
	var t int64

	for d, v := range l {

		fmt.Println(d, v)
		t += v
	}

	return t
}

func sumFloats(l map[string]float64) float64 {
	var t float64

	for d, v := range l {

		fmt.Println(d, v)
		t += v
	}

	return t
}

func sumIntsAndFloats[d comparable, v int64 | float64](l map[d]v) v {
	var t v

	for d, v := range l {
		fmt.Println(d, v)
		t += v
	}
	return t
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", sumInts(ints), sumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n", sumIntsAndFloats(ints), sumIntsAndFloats(floats))
}
