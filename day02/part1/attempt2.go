package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
Approach: Scanner

Benchmarks:
564 valid passwords!
Time: 0.001834035 seconds
 */

func main() {
	start := time.Now()
	file, err := os.Open("./../input.txt")
	if err != nil{
		fmt.Println(err, "\nYou need a file input.txt in the parent dir")
		return
	}
	defer file.Close()

	validPasswords := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var (
			lowerLimit, upperLimit int
			char rune
			password string
		)
		fmt.Sscanf(sc.Text(), "%d-%d %c: %s", &lowerLimit, &upperLimit, &char, &password) //Formatted scan of the input string

		if strings.Count(password, string(char)) >= lowerLimit && strings.Count(password, string(char)) <= upperLimit {
			validPasswords++
		}
	}
	fmt.Printf("%v valid passwords!\n", validPasswords)
	end := time.Now()
	fmt.Printf("Time: %v seconds\n", end.Sub(start).Seconds())
}
