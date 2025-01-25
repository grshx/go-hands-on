package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	add      = "+"
	subtract = "-"
	multiply = "*"
	divide   = "/"
	exponent = "**"
)

func main() {

	argLength := len(os.Args)

	if argLength != 4 {
		fmt.Println("invalid number of arguments, expected 3 got ", argLength-1)
		return

	}

	leftOperand, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("argument 1 is not a number")
		return
	}

	rightOperand, err := strconv.ParseFloat(os.Args[3], 64)

	if err != nil {
		fmt.Println("argument 3 is not a number")
		return
	}

	operation := os.Args[2]

	var result interface{}

	switch operation {

	case add:
		result = leftOperand + rightOperand
	case subtract:
		result = leftOperand - rightOperand
	case multiply:
		result = leftOperand * rightOperand
	case divide:
		result = leftOperand / rightOperand
	case exponent:
		result = math.Pow(leftOperand, rightOperand)
	default:
		fmt.Printf("oeration %s is not supported", operation)
		return
	}

	fmt.Printf("result: %v", result)
}
