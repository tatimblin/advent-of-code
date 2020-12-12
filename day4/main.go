package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	allPassports := []Passport{}
	currentPassport := Passport{}

	for _, line := range lines {
		if line != "" {
			entries := strings.Split(line, " ")
			r := make(map[string]string)
			for _, e := range entries {
				parts := strings.Split(e, ":")
				r[parts[0]] = parts[1]

				if val, ok := r["byr"]; ok {
					currentPassport.BirthYear = val
				}

				if val, ok := r["iyr"]; ok {
					currentPassport.IssueYear = val
				}

				if val, ok := r["eyr"]; ok {
					currentPassport.ExpirationYear = val
				}

				if val, ok := r["hgt"]; ok {
					currentPassport.Height = val
				}

				if val, ok := r["hcl"]; ok {
					currentPassport.HairColor = val
				}

				if val, ok := r["ecl"]; ok {
					currentPassport.EyeColor = val
				}

				if val, ok := r["pid"]; ok {
					currentPassport.PassportID = val
				}

				if val, ok := r["cid"]; ok {
					currentPassport.CountryID = val
				}
			}
		} else {
			allPassports = append(allPassports, currentPassport)
			currentPassport = Passport{}
		}
	}

	// Check for required fields
	valid := 0
	for _, e := range allPassports {
		if e.BirthYear != "" && e.ExpirationYear != "" && e.EyeColor != "" && e.HairColor != "" && e.Height != "" && e.IssueYear != "" && e.PassportID != "" {
			valid = valid + 1
		}
	}
	log.Println(valid)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
