package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sum int

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		l := s.Text()

		for r := range strings.SplitSeq(l, ",") {
			r = strings.TrimSpace(r)

			p := strings.Split(r, "-")

			if len(p) != 2 {
				continue
			}

			p1, err := strconv.Atoi(p[0])
			if err != nil {
				continue
			}

			p2, err := strconv.Atoi(p[1])
			if err != nil {
				continue
			}

			for i := p1; i < p2+1; i++ {
				s := strconv.Itoa(i)

				if len(s)%2 != 0 {
					continue
				}

				if s[0:len(s)/2] == s[len(s)/2:] {
					sum += i
					fmt.Println(i)
				}
			}
		}

		fmt.Println(sum)
	}

	if err := s.Err(); err != nil {
		fmt.Println("error scanning file")
		os.Exit(1)
	}

}
