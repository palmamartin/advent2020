package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var count int
	p := new(Passport)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		p.ParseLine(line)

		// New passport starts.
		if line == "" {
			// Check if prev. passport is valid.
			if p.IsValid() {
				count++
			}
			p = new(Passport)
		}
	}

	if p.IsValid() {
		count++
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func height(hgt string) string {
	var hgtRe = regexp.MustCompile(`^(\d*)(cm|in)$`)

	f := hgtRe.FindStringSubmatch(hgt)
	if len(f) != 3 {
		return ""
	}
	h, err := strconv.Atoi(f[1])
	if err != nil {
		return ""
	}

	switch f[2] {
	case "cm":
		if h < 150 || h > 193 {
			return ""
		}
	case "in":
		if h < 59 || h > 76 {
			return ""
		}
	}
	return f[0]
}

type Passport struct {
	byr, iyr, eyr           int
	hgt, hcl, ecl, pid, cid string
}

func (p *Passport) ParseLine(line string) error {
	var (
		err   error
		hclRe = regexp.MustCompile(`^#([0-9]|[a-f]){6}$`)
		eclRe = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		pidRe = regexp.MustCompile(`^\d{9}$`)
	)

	for _, f := range strings.Fields(strings.TrimSpace(line)) {
		s := strings.Split(f, ":")
		if len(s) != 2 {
			return errors.New("parsing error")
		}

		k, v := s[0], s[1]
		switch k {
		case "byr":
			p.byr, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		case "iyr":
			p.iyr, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		case "eyr":
			p.eyr, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		case "hgt":
			p.hgt = height(v)
		case "hcl":
			p.hcl = hclRe.FindString(v)
		case "ecl":
			p.ecl = eclRe.FindString(v)
		case "pid":
			p.pid = pidRe.FindString(v)
		case "cid":
			p.cid = v
		}
	}

	return nil
}
func (p *Passport) IsValid() bool {
	switch {
	case p.byr < 1920 || p.byr > 2002:
		return false
	case p.iyr < 2010 || p.iyr > 2020:
		return false
	case p.eyr < 2020 || p.eyr > 2030:
		return false
	case p.hgt == "":
		return false
	case p.hcl == "":
		return false
	case p.ecl == "":
		return false
	case p.pid == "":
		return false
	default:
		return true
	}
}
