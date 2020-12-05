package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
/*
Approach: Original
Valid Passports: 188
Time: 0.000893911 seconds
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
		key := splitCred[0]
		credsMap[key] = true
		data := splitCred[1]

		valid := true

		switch key {
		case "byr":
			valid = validateRange(data, 1920, 2002)
		case "iyr":
			valid = validateRange(data, 2010, 2020)
		case "eyr":
			valid = validateRange(data, 2020, 2030)
		case "hgt":
			matches := hgtRegexp.FindStringSubmatch(data)
			if len(matches) != 3 {
				valid = false
			} else {
				num := matches[1]
				if matches[2] == "cm" {
					valid = validateRange(num, 150, 193)
				} else if matches[2] == "in" {
					valid = validateRange(num, 59, 76)
				}
			}
		case "hcl":
			valid = hclRegexp.MatchString(data)
		case "ecl":
			_, valid = eyecolors[data]
		case "pid":
			valid = pidRegexp.MatchString(data)
		}
		if !valid {
			return false
		}
	}

	for _, it := range required {
		_, ok := credsMap[it]
		if !ok {
			return false
		}
	}
	return true
}

var eyecolors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

var pidRegexp = regexp.MustCompile("^[0-9]{9}$")
var hclRegexp = regexp.MustCompile("^#[0-9a-f]{6}$")
var hgtRegexp = regexp.MustCompile(`^(\d{2,3})(\w{2})$`)

func validateRange(val string, min, max int) bool {
	n, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if n >= min && n <= max {
		return true
	}
	return false
}