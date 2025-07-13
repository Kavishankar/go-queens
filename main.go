package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	inData, err := os.ReadFile("./data/input.txt")
	if err != nil {
		log.Panic(err)
	}
	outData, err := os.ReadFile("./data/output.txt")
	if err != nil {
		log.Panic(err)
	}
	g, err := NewGrid(inData)
	if err != nil {
		log.Panic(err)
	}
	g.Print()
	if !g.Solve() {
		log.Panic("Grid is not solved. Plzzz Help!")
	}
	if !g.Validate(outData) {
		log.Panic("The program's solution is not valid :(")
	} else {
		log.Println("Your solution is correct!")
	}
}

func ValidateInput(input []byte) (int, error) {
	// TODO: Implement ValidateInput

	length := len(input)
	dim := int(math.Sqrt(float64(length)))

	if dim*dim != length {
		return 0, errors.New("input is not a perfect square")
	}

	maxAllowedChar := byte('a' + dim - 1)

	for i, b := range input {
		if b < 'a' || b > maxAllowedChar {
			return 0, fmt.Errorf("invalid character '%c' at position %d: must be between 'a' and '%c'", b, i, maxAllowedChar)
		}
	}

	return dim, nil
}
