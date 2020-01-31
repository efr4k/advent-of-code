package main

import (
	"bufio"
	"fmt"
	"os"
)

var fuel = 0

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("could not open file; got %v", err))
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var mass int
		_, err = fmt.Sscanf(s.Text(), "%d", &mass)
		if err != nil {
			panic(fmt.Sprintf("could not Sscanf %s; got %v", s.Text(), err))
		}
		addFuel(mass)
		//fuel += (mass/3) - 2
	}
	if err = s.Err(); err != nil {
		panic(fmt.Sprintf("scanner error; got %v", err))
	}
	fmt.Println(fuel)
}

func addFuel(mass int) {
	potentialFuel := (mass/3) - 2
	if potentialFuel > 0 {
		fuel += potentialFuel
		addFuel(potentialFuel)
	}
}