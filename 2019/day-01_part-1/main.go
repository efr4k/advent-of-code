package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("could not open file; got %v", err))
	}
	defer f.Close()

	fuel := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		var mass int
		_, err = fmt.Sscanf(s.Text(), "%d", &mass)
		if err != nil {
			panic(fmt.Sprintf("could not Sscanf %s; got %v", s.Text(), err))
		}
		fuel += (mass/3) - 2
	}
	if err = s.Err(); err != nil {
		panic(fmt.Sprintf("scanner error; got %v", err))
	}
	fmt.Println(fuel)
}
