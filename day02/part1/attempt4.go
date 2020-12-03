package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
Approach: Regex

Benchmarks:
Valid Passwords: 564
Time: 0.007633162 seconds
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
		regex := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
		matches := regex.FindStringSubmatch(l)
		if len(matches) != 5 {
			fmt.Println("Bad line")
			continue
		}
		min, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		count := strings.Count(matches[4], matches[3])
		if count >= min && count <= max {
			validPwCount++
		}
	}

	fmt.Printf("Valid Passwords: %v", validPwCount)
	end := time.Now()
	fmt.Printf("Time: %v seconds", end.Sub(start).Seconds())
}

