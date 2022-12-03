package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func validate(passport map[string]string) bool {
	for key, value := range passport {

		if !validateKey(key, value) {
			fmt.Println(key, value)
			return false
		}
	}

	return true
}

func arrayContains(needle string, haystack []string) bool {
	for _, x := range haystack {
		if x == needle {
			return true
		}
	}
	return false
}

func validateKey(key string, value string) bool {
	if key == "byr" {
		num, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if num < 1920 || num > 2002 {
			return false
		}
	}

	if key == "iyr" {
		num, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if num < 2010 || num > 2020 {
			return false
		}
	}

	if key == "eyr" {
		num, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if num < 2020 || num > 2030 {
			return false
		}
	}

	eyecolors := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	if key == "ecl" {
		if !arrayContains(value, eyecolors[:]) {
			return false
		}
	}

	if key == "pid" {
		if len(value) != 9 {
			return false
		}
		_, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
	}

	if key == "hcl" {
		rexp, _ := regexp.Compile("#[0-9a-f]{6}")
		match := rexp.MatchString(value)
		if !match {
			return false
		}
	}

	if key == "hgt" {
		last2 := value[len(value)-2:]
		if last2 != "cm" && last2 != "in" {
			return false
		}
		if last2 == "cm" {
			num, err := strconv.Atoi(value[0:3])
			if err != nil {
				return false
			}
			if num < 150 || num > 193 {
				return false
			}
		}
		if last2 == "in" {
			num, err := strconv.Atoi(value[0:2])
			if err != nil {
				return false
			}
			if num < 59 || num > 76 {
				return false
			}
		}
	}

	return true
}

func main() {
	cont, err := os.ReadFile("./data.txt")
	if err != nil {
		fmt.Print(err)
	}

	contStr := string(cont)

	passports := strings.Split(contStr, "\n\n")
	task1 := 0
	task2 := 0

	for _, passport := range passports {
		values := strings.Fields(passport)
		current := make(map[string]string)
		for _, kvpair := range values {
			split := strings.Split(kvpair, ":")
			key, value := split[0], split[1]
			current[key] = value
		}
		_, ok := current["cid"]
		if (len(current) == 8) || (len(current) == 7 && !ok) {
			task1++
			if validate(current) {
				task2++
			}
		}
	}
	fmt.Printf("Task 1: %d\n", task1)
	fmt.Printf("Task 2: %d\n", task2)
}
