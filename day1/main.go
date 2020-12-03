package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	exp := make(map[int]struct{})

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		exp[i] = struct{}{}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for k, _ := range exp {
		d := 2020 - k

		for k1, _ := range exp {
			n := d - k1
			_, ok := exp[n]
			if !ok {
				continue
			}
			fmt.Println(k * n * k1)
			return
		}
	}

}