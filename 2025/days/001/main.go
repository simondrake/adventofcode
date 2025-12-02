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

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	dial := 50

	var count int

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

		if dial == 0 {
			count++
		}
	}

	fmt.Println(count)
}

func processRotation(dial int, rotation int, direction byte) int {
	var n int

	switch direction {
	case 76:
		n = dial - rotation
		for n < 0 {
			n = 100 - n*-1
		}
	case 82:
		n = dial + rotation

		for n >= 100 {
			n = 0 + (n - 100)
		}
	}

	return n
}
