package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

/*
Approach: One Coordinate At A Time
Benchmark:
Total Tree Multiplication: 3517401300
Time: 0.000930631 seconds
*/

var incrementers = [][]int{
	{ 1, 1},
	{ 3, 1},
	{ 5, 1},
	{ 7, 1},
	{ 1, 2},
}

func main() {
	start := time.Now()
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file input.txt in the parent dir")
		return
	}

	lines := strings.Split(string(b), "\n")
	treeMult := 1
	for _, row := range incrementers {
		treeMult *= findTreeCount(lines, row[0], row[1])
	}
	fmt.Printf("Total Tree Multiplication: %v\n", treeMult)
	end := time.Now()
	fmt.Printf("Time: %v seconds\n", end.Sub(start).Seconds())
}



func findTreeCount(lines []string, xIncrementer, yIncrementer int) (treeCount int) {
	x := 0
	repeatCounter := 1
	for i := 0; i < len(lines); {
		// Empty line occurs at the end of the file when we use Split.
		l := lines[i]
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
		x = x + xIncrementer // move across our x counter
		i = i + yIncrementer // move down rows
	}

	fmt.Printf("Tree Count: %v\n", treeCount)
	return
}

