package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

/*
Approach: Original
Valid Passports: 239
Time: 0.000602789 seconds
*/
func main()  {
	start := time.Now()
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file input.txt in the parent dir")
		return
	}

	lines := strings.Split(string(b), "\n")
	validPassports := 0
	var captureData string
	for _, l := range lines {
		if len(l) == 0 {
			if isValid(captureData) {
				validPassports++
			}
			captureData = ""
			continue
		}
		captureData = strings.Join([]string{captureData, l}, " ")
	}
	if len(captureData) != 0 && isValid(captureData){
		validPassports++
	}

	fmt.Printf("Valid Passports: %v\n", validPassports)
	end := time.Now()
	fmt.Printf("Time: %v seconds\n", end.Sub(start).Seconds())
}

func isValid(captureData string) (valid bool) {
	trimmed := strings.TrimPrefix(captureData, " ") // space starts captureData bc of Join
	creds := strings.Split(trimmed, " ")
	credsMap := map[string]bool{}
	for _, cred := range creds {
		splitCred := strings.Split(cred, ":")
		credsMap[splitCred[0]] = true
	}
	invalid := false
	for _, it := range required {
		_, ok := credsMap[it]
		if !ok {
			invalid = true
			break
		}
	}
	return !invalid
}