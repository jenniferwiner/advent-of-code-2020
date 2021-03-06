package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

/*
Approach: Scanner
Benchmarks:
0.001800296 seconds
 */
func main() {
	start := time.Now()
	file, err := os.Open("../input.txt")
	if err != nil{
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	var validPasswords int

	sc := bufio.NewScanner(file)
	for(sc.Scan()){
		var (
			lowerPosition, upperPosition int
			char rune
			password string
		)
		fmt.Sscanf(sc.Text(), "%d-%d %c: %s", &lowerPosition, &upperPosition, &char, &password) //Formatted scan of the input string

		if (rune(password[lowerPosition-1]) == char) != (rune(password[upperPosition-1]) == char) { //Xor between two boolean values
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
	end := time.Now()
	fmt.Printf("Time: %v seconds", end.Sub(start).Seconds())
}
