package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madsaune/advent-of-code/internal/utils"
)

func main() {
	var input string
	var days int

	flag.StringVar(&input, "input", "input.txt", "Path to input file")
	flag.IntVar(&days, "days", 80, "number of days")
	flag.Parse()

	lines, err := utils.Readlines(input)
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	// -----------------------------------------------------------------------------
	//     - CONVERT INPUT TO INT -
	// -----------------------------------------------------------------------------
	parsed := strings.Split(lines[0], ",")
	fishes := make([]int, len(parsed))
	for i, v := range parsed {
		n, _ := strconv.Atoi(v)
		fishes[i] = n
	}

	// -----------------------------------------------------------------------------
	//     - SOlUTION 1 -
	//
	// Worked fine for part 1, but to slow on part 2
	// -----------------------------------------------------------------------------
	// for i := 0; i < days; i += 1 {
	// 	for idx, f := range fishes {
	// 		if f == 0 {
	// 			fishes[idx] = 6
	// 			fishes = append(fishes, 8)
	// 		} else {
	// 			fishes[idx] -= 1
	// 		}
	// 	}
	// }

	// -----------------------------------------------------------------------------
	//     - SOLUTION 2 -
	//
	// Keep track of each state instead of each fish
	// credit: https://barretblake.dev/blog/2021/12/advent-of-code-day6/
	// -----------------------------------------------------------------------------
	var fishStates [9]int
	for _, f := range fishes {
		fishStates[f] += 1
	}

	for i := 0; i < days; i += 1 {
		// store count of fishes that will reproduce this day
		reproducingFishes := fishStates[0]

		// shift all fishes one day down
		for j := 0; j < 8; j += 1 {
			fishStates[j] = fishStates[j+1]
		}

		// reset reproducing fishes
		fishStates[6] += reproducingFishes

		// the amount of new fishes is the same as
		// reproducing fishes
		fishStates[8] = reproducingFishes
	}

	var fishCount int
	for x := 0; x <= 8; x += 1 {
		fishCount += fishStates[x]
	}

	fmt.Printf("How many lanternfish would there be after %d days? ==> %d\n", days, fishCount)
}
