package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	exit := false
	input, err := ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	for _, first := range input {
		secondaryTarget := 2020 - first
		// seed map with everything but `first`
		entriesMinusKey := map[int]bool{}
		for _, val := range input {
			if first != val {
				entriesMinusKey[val] = true
			}
		}

		// loop over map
		for second, _ := range entriesMinusKey {
			// for each entry, looking for third value
			third := secondaryTarget - second
			if _, ok := entriesMinusKey[third]; ok {
				fmt.Printf("Found them: %v, %v, %v\n", first,  second, third)
				fmt.Printf("Multiply: %v\n", first * second * third)
				exit = true
				break
			}
		}
		if exit {
			break
		}
	}
	end := time.Now()
	fmt.Printf("Time: %v seconds", end.Sub(start).Seconds())
}

func readFile(filename string) (nums []int, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
