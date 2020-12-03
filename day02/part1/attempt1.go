package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/*
Approach: Structs and All - Original Sol'n

Benchmarks:
Valid Passwords: 564
Time: 0.000649219 seconds
 */
func main() {
	start := time.Now()
	input, err := buildMinMaxPWMap("../input.txt")
	if err != nil {
		panic(err)
	}
	validPwCount := 0
	for pw, entry := range input {
		count := strings.Count(pw, entry.letter)
		if count >= entry.min && count <= entry.max {
			validPwCount++
		}
	}
	fmt.Printf("Valid Passwords: %v", validPwCount)
	end := time.Now()
	fmt.Printf("Time: %v seconds", end.Sub(start).Seconds())
}

type minMaxPWDetails struct {
	letter string
	min int
	max int
}

func buildMinMaxPWMap(filename string) (map[string]minMaxPWDetails, error) {
	path, err := filepath.Abs(filename)
	fmt.Println(path)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	entries := map[string]minMaxPWDetails{}
	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		contents := strings.Split(l, " ")
		numStr := contents[0]
		nums := strings.Split(numStr, "-")
		let := contents[1]
		pw := contents[2]

		min, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}
		max, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, err
		}
		entries[pw] = minMaxPWDetails{
			letter: let[:1],
			min:    min,
			max:    max,
		}
	}

	return entries, nil
}

