package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/madsaune/advent-of-code/internal/utils"
)

type position struct {
	x   int
	y   int
	aim int
}

func taskOne(f *os.File) {
	defer f.Close()

	p := position{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		cmd := s[0]
		unit, _ := strconv.Atoi(s[1])

		switch cmd {
		case "forward":
			p.x += unit
		case "up":
			p.y -= unit
		case "down":
			p.y += unit
		}
	}

	answer := p.x * p.y

	fmt.Printf("What do you get if you multiply your final horizontal position by your final depth? ==> %d\n", answer)
}

func taskTwo(f *os.File) {
	defer f.Close()

	p := position{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		cmd := s[0]
		unit, _ := strconv.Atoi(s[1])

		switch cmd {
		case "forward":
			p.x += unit
			p.y += p.aim * unit
		case "up":
			p.aim -= unit
		case "down":
			p.aim += unit
		}
	}

	answer := p.x * p.y
	fmt.Printf("What do you get if you multiply your final horizontal position by your final depth? ==> %d\n", answer)
}

func main() {
	var input1 string
	var input2 string

	flag.StringVar(&input1, "input1", "input1.txt", "Path to input file for task 1")
	flag.StringVar(&input2, "input2", "input2.txt", "Path to input file for task 2")
	flag.Parse()

	taskOne(utils.LoadFile(input1))
	taskTwo(utils.LoadFile(input2))
}
