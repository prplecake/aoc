package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("File reading error:", err)
		return
	}
	input := string(file)
	inputSlice := parse(input)

	cases := []*regexp.Regexp{
		regexp.MustCompile(`pid:([^ ]+)`),
		regexp.MustCompile(`iyr:([^ ]+)`),
		regexp.MustCompile(`ecl:([^ ]+)`),
		regexp.MustCompile(`byr:([^ ]+)`),
		regexp.MustCompile(`eyr:([^ ]+)`),
		regexp.MustCompile(`hcl:([^ ]+)`),
		regexp.MustCompile(`hgt:([^ ]+)`),
	}

	var validPassports int
	for _, v := range inputSlice {
		valid := true
		for _, c := range cases {
			if !c.MatchString(v) {
				valid = false
				break
			}
		}
		if valid {
			validPassports++
		}
	}
	fmt.Println("Valid passports (p1):", validPassports)

	var passports []Passport
	for _, x := range inputSlice {
		passports = append(passports, NewPassport(x))
	}

	validPassports = 0
	for _, passport := range passports {
		if passport.IsValid() {
			validPassports++
		}
	}
	fmt.Println("Valid passports (p2):", validPassports)
}

func parse(input string) []string {
	inputSlice := strings.Split(strings.TrimSpace(input), "\n\n")
	for i, x := range inputSlice {
		inputSlice[i] = strings.ReplaceAll(x, "\n", " ")
	}
	return inputSlice
}

type Passport struct {
	byr, iyr, cid, pid, hgt, hcl, ecl, eyr string
}

func (p Passport) IsValid() bool {
	if p.byr == "" {
		return false
	}
	{
		i, _ := strconv.Atoi(p.byr)
		if !(1920 <= i && i <= 2002) {
			return false
		}
	}

	if p.iyr == "" {
		return false
	}
	{
		i, _ := strconv.Atoi(p.iyr)
		if !(2010 <= i && i <= 2020) {
			return false
		}
	}
	if len(p.pid) == 9 {
		_, err := strconv.Atoi(p.pid)
		if err != nil {
			return false
		}
	} else {
		return false
	}

	if p.hgt == "" {
		return false
	}
	{
		if strings.Contains(p.hgt, "cm") {
			i, _ := strconv.Atoi(strings.Trim(p.hgt, "cm"))
			if !(150 <= i && i <= 193) {
				return false
			}
		} else if strings.Contains(p.hgt, "in") {
			i, _ := strconv.Atoi(strings.Trim(p.hgt, "in"))
			if !(59 <= i && i <= 76) {
				return false
			}
		} else {
			return false
		}
	}

	if p.hcl == "" {
		return false
	}
	{
		if p.hcl[0] == []byte("#")[0] {
			if m, _ := regexp.MatchString(`#[0-9a-f]{6}`, p.hcl); !m {
				return false
			}
		} else {
			return false
		}
	}

	if p.ecl == "" {
		return false
	}
	{
		var contained bool
		for _, v := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if p.ecl == v {
				contained = true
				break
			}
		}
		if !contained {
			return false
		}
	}

	if p.eyr == "" {
		return false
	}
	{
		i, _ := strconv.Atoi(p.eyr)
		if !(2020 <= i && i <= 2030) {
			return false
		}
	}

	return true
}

func NewPassport(input string) Passport {
	np := Passport{}
	np.byr = extractField("byr", input)
	np.iyr = extractField("iyr", input)
	np.cid = extractField("cid", input)
	np.pid = extractField("pid", input)
	np.hgt = extractField("hgt", input)
	np.hcl = extractField("hcl", input)
	np.ecl = extractField("ecl", input)
	np.eyr = extractField("eyr", input)
	return np
}

func extractField(field, input string) string {
	fieldRegex := regexp.MustCompile(field + `:([^ ]+)`)
	matches := fieldRegex.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		return ""
	}
	return matches[0][1]
}
