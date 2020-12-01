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
	input, err := ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	hashmap := map[int]bool{}
	for _, val := range input {
		lookingFor := 2020 - val
		_, ok := hashmap[lookingFor]
		if ok {
			fmt.Printf("Found them: %v, %v\n", val, lookingFor)
			fmt.Printf("Multiply: %v\n", val*lookingFor)
			break
		}
		hashmap[val] = true
	}
	end := time.Now()
	fmt.Printf("Time: %v seconds", end.Sub(start).Seconds())
}

func ReadFile(filename string) (nums []int, err error) {
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

