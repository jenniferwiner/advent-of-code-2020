package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

/*
Approach: Remove Structs

Benchmarks:
Valid Passwords: 564
Time: 0.000356604 seconds
 */
func main() {
	start := time.Now()

	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file input.txt in the parent dir")
		return
	}

	validPwCount := 0
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		contents := strings.Split(l, " ")
		numbers := strings.Split(contents[0], "-")
		letter := contents[1][:1]
		password := contents[2]

		min, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		count := strings.Count(password, letter)
		if count >= min && count <= max {
			validPwCount++
		}
	}

	fmt.Printf("Valid Passwords: %v", validPwCount)
	end := time.Now()
	fmt.Printf("Time: %v seconds", end.Sub(start).Seconds())
}

