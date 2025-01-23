package main

import (
	"fmt"
	"time"
)

func FizzBuzz() {
	start := time.Now()
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, " - FizzBuzz!")
			continue
		}

		if i%3 == 0 {
			fmt.Println(i, " - Fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println(i, " - Buzz")
			continue
		}

		fmt.Println(i)

	}
	fmt.Println("the program took ", time.Since(start))
}

func main() {
	FizzBuzz()
}
