package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)
/*
Approach: Binary
Part 1: Highest Seat ID: 861
Part 2: Found Seat ID: 633
Time: 0.000711054 seconds
*/

func main()  {
	start := time.Now()
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file input.txt in the curr dir")
		return
	}
	highestSeatID := 0
	lowestSeatID := 1000
	lines := strings.Split(string(b), "\n")
	seatsMap := map[int]bool{}
	for _, l := range lines {
		seatID, err := findSeatID(l)
		if err != nil {
			fmt.Println(err)
			return
		}
		seatsMap[seatID] = true
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
		if seatID < lowestSeatID {
			lowestSeatID = seatID
		}
	}
	var foundSeat int
	for i := lowestSeatID; i < highestSeatID + 1; i++ {
		if _, ok := seatsMap[i]; !ok {
			foundSeat = i
			break
		}
	}
	fmt.Printf("Part 1: Highest Seat ID: %v\n", highestSeatID)
	fmt.Printf("Part 2: Found Seat ID: %v\n", foundSeat)
	end := time.Now()
	fmt.Printf("Time: %v seconds\n", end.Sub(start).Seconds())
}

func findSeatID(l string) (int, error) {
	row := ""
	column := ""
	for _, char := range l {
		letter := string(char)
		switch letter {
		case "F" :
			row += "0"
		case "B":
			row += "1"
		case "L":
			column += "0"
		case "R":
			column += "1"
		}
	}
	rowVal, err := strconv.ParseInt(row, 2, 64)
	if err != nil {
		return 0, err
	}
	columnVal, err := strconv.ParseInt(column, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(rowVal*8+columnVal), nil

}