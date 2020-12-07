package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

/*
Benchmarks
Question Count: 3288
Time: 0.001081733 seconds
 */
func main() {
	start := time.Now()
	file, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file input.txt in the curr dir")
		return
	}
	grp := strings.Split(string(file), "\n\n")
	questionCount := 0
	for _, g := range grp {
		questionMap := map[rune]int{}
		lines := strings.Split(g, "\n")
		for _, l := range lines {
			for _, char := range l {
				questionMap[char]++
			}
		}
		for _, val := range questionMap {
			if val == len(lines) {
				questionCount++
			}
		}
	}
	fmt.Printf("Question Count: %v\n", questionCount)
	end := time.Now()
	fmt.Printf("Time: %v seconds\n", end.Sub(start).Seconds())
}
