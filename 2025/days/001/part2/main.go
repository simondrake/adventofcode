package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var count int

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	dial := 50

	for {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("error reading from file: %v\n", err)
			os.Exit(1)
		}

		l = strings.TrimSpace(l)

		if errors.Is(err, io.EOF) || len(l) < 2 {
			break
		}

		d := l[0]
		l = l[1:]

		i, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("error converting to int: %v\n", err)
			os.Exit(1)
		}

		dial = processRotation(dial, i, d)
	}

	fmt.Println(count)
}

func processRotation(dial int, rotation int, direction byte) int {

	var n int

	switch direction {
	case 'L':
		n = dial - rotation

		// we need the wrapped var so we can keep track of situations where
		// a large step (e.g. L1000) causes it to loop round/wrap multiple times,
		// as we don't want to lose those instances
		wrapped := dial != 0

		for n <= 0 {
			n = 100 - -n

			if wrapped {
				count++
			}

			wrapped = true

		}

		if n == 100 {
			n = 0
		}

	case 'R':
		n = dial + rotation

		for n >= 100 {
			count++
			n = n - 100
		}
	}

	return n
}
