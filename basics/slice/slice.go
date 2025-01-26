package main

import "fmt"

func SliceOps() {
	var slice []int

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)

	fmt.Println("slice : ", slice)
	fmt.Println("slice length: ", len(slice))
	fmt.Println("capacity : ", cap(slice))
}

func SliceTheSlice() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	subSlice := slice[2:5]

	fmt.Println(subSlice)

}

func ModifySlice() {
	slice := []int{13, 45, 25, -35, 756, 35}
	slice[2] = -22
	fmt.Println(slice)

	// delete element from slice
	slice = append(slice[:3], slice[4:]...)
	fmt.Println(slice)
}

func main() {
	ModifySlice()
}
