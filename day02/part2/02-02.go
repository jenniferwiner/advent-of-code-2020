package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type posPWDetails struct {
	letter string
	pos []int
}

func main() {
	input, err := buildPosPWMap("../input.txt")
	if err != nil {
		panic(err)
	}
	validPwCount := 0
	for pw, entry := range input {
		count := 0
		for p, v := range pw {
			currLetter := string(v)
			if entry.pos[0] == p && currLetter == entry.letter {
				count++
			}
			if entry.pos[1] == p && currLetter == entry.letter {
				count++
			}
		}
		if count == 1 {
			validPwCount++
		}
	}
	fmt.Printf("Valid Passwords: %v", validPwCount)
}

func buildPosPWMap(filename string) (map[string]posPWDetails, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	entries := map[string]posPWDetails{}
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

		pos1, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}
		pos2, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, err
		}
		entries[pw] = posPWDetails{
			letter: let[:1],
			// normalize indexes
			pos: []int{pos1 - 1, pos2 - 1},
		}
	}

	return entries, nil
}

