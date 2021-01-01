package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	pc := parse(input)

	fmt.Println(pc.validCount())
	fmt.Println(pc.strictValidCount())
}

func parse(in string) collection {
	pc := collection{}
	p := passport{}

	r := regexp.MustCompile(`([a-z]{3}):([A-Za-z0-9#]+)`)

	for _, d := range strings.Split(in, "\n") {
		if len(d) < 1 {
			pc = append(pc, p)
			p = passport{}
			continue
		}

		m := r.FindAllStringSubmatch(d, -1)

		for _, md := range m {
			if len(md) != 3 {
				continue
			}

			k, v := md[1], md[2]

			switch k {
			case "byr":
				p.byr = v
			case "iyr":
				p.iyr = v
			case "eyr":
				p.eyr = v
			case "hgt":
				p.hgt = v
			case "hcl":
				p.hcl = v
			case "ecl":
				p.ecl = v
			case "pid":
				p.pid = v
			case "cid":
				p.cid = v
			}
		}
	}

	return append(pc, p) // make sure last is appended
}

type collection []passport

func (pc collection) strictValidCount() int {
	count := 0
	for _, p := range pc {
		if p.strictValid() {
			count++
		}
	}
	return count
}

func (pc collection) validCount() int {
	count := 0
	for _, p := range pc {
		if p.valid() {
			count++
		}
	}
	return count
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p passport) strictValid() bool {
	yr := regexp.MustCompile(`^\d{4}$`)
	hgt := regexp.MustCompile(`^\d+(cm|in)$`)
	hcl := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	ecl := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pid := regexp.MustCompile(`^\d{9}$`)

	return p.valid() &&
		yr.MatchString(p.byr) &&
		p.byr >= "1920" &&
		p.byr <= "2002" &&
		yr.MatchString(p.iyr) &&
		p.iyr >= "2010" &&
		p.iyr <= "2020" &&
		yr.MatchString(p.eyr) &&
		p.eyr >= "2020" &&
		p.eyr <= "2030" &&
		hgt.MatchString(p.hgt) &&
		((strings.HasSuffix(p.hgt, "cm") && p.hgt >= "150cm" && p.hgt <= "193cm") ||
			(strings.HasSuffix(p.hgt, "in") && p.hgt >= "59in" && p.hgt <= "76in")) &&
		hcl.MatchString(p.hcl) &&
		ecl.MatchString(p.ecl) &&
		pid.MatchString(p.pid)
}

func (p passport) valid() bool {
	return len(p.byr) > 0 &&
		len(p.iyr) > 0 &&
		len(p.eyr) > 0 &&
		len(p.hgt) > 0 &&
		len(p.hcl) > 0 &&
		len(p.ecl) > 0 &&
		len(p.pid) > 0
}
