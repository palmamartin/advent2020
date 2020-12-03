package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var part1, part2 int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		edges := strings.Split(fields[0], "-")
		char := strings.TrimSuffix(fields[1], ":")
		password := fields[2]
		lower, err := strconv.Atoi(edges[0])
		if err != nil {
			log.Fatal(err)
		}

		upper, err := strconv.Atoi(edges[1])
		if err != nil {
			log.Fatal(err)
		}

		count := strings.Count(password, char)
		if count >= lower && count <= upper {
			part1++
		}

		if (password[lower-1] == char[0]) != (password[upper-1] == char[0]) {
			part2++
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
