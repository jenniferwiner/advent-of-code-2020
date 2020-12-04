package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

/*
Approach: Original
Benchmark:
Tree Count: 223
Time: 0.000284215 seconds
 */
func main() {
	start := time.Now()
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file input.txt in the parent dir")
		return
	}

	lines := strings.Split(string(b), "\n")
	treeCount := 0
	x := 0
	repeatCounter := 1
	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		if x >= len(l) * repeatCounter {
			repeatCounter++
		}
		lineToEval := strings.Repeat(l, repeatCounter)
		val := rune(lineToEval[x])
		if string(val) == "#" {
			treeCount++
		}
		x = x + 3 // move across our x counter
	}

	fmt.Printf("Tree Count: %v\n", treeCount)
	end := time.Now()
	fmt.Printf("Time: %v seconds\n", end.Sub(start).Seconds())
}
