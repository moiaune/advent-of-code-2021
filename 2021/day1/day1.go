package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/madsaune/advent-of-code/internal/utils"
)

func main() {

	var input1 string
	var input2 string

	flag.StringVar(&input1, "input1", "input1.txt", "Path to input file for task 1")
	flag.StringVar(&input2, "input2", "input2.txt", "Path to input file for task 2")
	flag.Parse()

	taskOne(utils.LoadFile(input1))
	taskTwo(utils.LoadFile(input2))
}

type previousDepth struct {
	value int
	valid bool
}

func taskOne(f *os.File) {
	defer f.Close()

	var count int
	prev := previousDepth{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("ERR: could not convert depth(str) to int: %v", err)
		}

		if !prev.valid {
			prev.value = curr
			prev.valid = true
			continue
		}

		if prev.value <= curr {
			count += 1
		}

		prev.value = curr
	}

	fmt.Printf("Task 1: How many measurements are larger than the previous measurement? ==> %d\n", count)
}

func taskTwo(f *os.File) {
	defer f.Close()

	var count int
	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for idx := range lines {
		if (idx + 3) <= (len(lines) - 1) {
			a, _ := strconv.Atoi(lines[idx])
			b, _ := strconv.Atoi(lines[idx+1])
			c, _ := strconv.Atoi(lines[idx+2])
			d, _ := strconv.Atoi(lines[idx+3])

			windowA := a + b + c
			windowB := b + c + d

			if windowA < windowB {
				count += 1
			}
		}
	}

	fmt.Printf("Task 2: How many measurements are larger than the previous measurement? ==> %d\n", count)
}
