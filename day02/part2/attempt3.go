package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

/*
Approach: Runes
Benchmarks:
0.000303921 seconds
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
		letter := []rune(contents[1])[0]
		password := contents[2]

		upperLimit, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		lowerLimit, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		if (rune(password[lowerLimit-1]) == letter) != (rune(password[upperLimit-1]) == letter) { //Xor between two boolean values
			validPwCount++
		}
	}

	fmt.Printf("Valid Passwords: %v", validPwCount)
	end := time.Now()
	fmt.Printf("Time: %v seconds", end.Sub(start).Seconds())
}