package main

import "fmt"

func main() {
	//declare and print
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(numbers["two"])

	// check for key
	idx, ok := numbers["four"]
	fmt.Println("value exists :", ok, " at index:", idx)

	// iterate
	for k, v := range numbers {
		fmt.Println("key: ", k, "value: ", v)
	}
}
