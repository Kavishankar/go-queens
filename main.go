package main

import (
	"log"
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
	return 9, nil
}
