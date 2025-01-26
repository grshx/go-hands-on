package main

import (
	"fmt"
	"sync"
)

func sumRange(start, end int, res chan int, wg *sync.WaitGroup) {
	var sum int
	for i := start; i < end+1; i++ {
		sum += i
	}
	fmt.Printf("sum of number %d - %d: %d\n", start, end, sum)
	res <- sum
	wg.Done()
}

func main() {

	result := make(chan int, 3)

	var wg sync.WaitGroup

	for i := 1; i < 31; i = i + 10 {
		wg.Add(1)
		go sumRange(i, i+9, result, &wg)
	}
	wg.Wait()

	close(result)

	var totalSum int
	for sum := range result {
		totalSum += sum
	}

	fmt.Println("total sum : ", totalSum)
}
