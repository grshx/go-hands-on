package main

import "fmt"

var numberToWord = map[int32]string{
	1: "one",
	2: "two",
}

func main() {
	fmt.Printf("%v", numberToWord)
}
