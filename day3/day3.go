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
	var bitLength int

	flag.StringVar(&input, "input", "input.txt", "Path to input file")
	flag.IntVar(&bitLength, "bitlength", 12, "Length of bit in input")
	flag.Parse()

	counts := make([]int, bitLength)
	lines, err := utils.Readlines(input)
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	for _, line := range lines {
		s := strings.Split(line, "")

		// check how many times "1" occurs pr index
		for i, c := range s {
			if c == "1" {
				counts[i] += 1
			}
		}
	}

	var gamma string
	var epsilon string
	for idx := 0; idx < len(lines[0]); idx += 1 {
		// if the "1" at index is greater than half the lines
		// it must be the dominating value, else its "0"
		if counts[idx] > len(lines)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gd, _ := strconv.ParseInt(gamma, 2, 64)
	ed, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("What is the power consumption of the submarine? ==> %d\n", gd*ed)

	// -----------------------------------------------------------------------------
	//     - OXYGEN RATING -
	// -----------------------------------------------------------------------------
	oxygenRating := make([]string, len(lines))
	copy(oxygenRating, lines)

	for i := 0; i < len(gamma); i += 1 {

		var ones float64
		for _, l := range oxygenRating {
			if l[i] == '1' {
				ones += 1
			}
		}

		halfOfOxygen := float64(len(oxygenRating)) / 2

		// Oxygen Generator
		for j := len(oxygenRating) - 1; j >= 0; j -= 1 {
			if ones >= halfOfOxygen {
				if oxygenRating[j][i] == '0' {
					oxygenRating = utils.RemoveFromSlice(oxygenRating, j)
				}
			} else {
				if oxygenRating[j][i] == '1' {
					oxygenRating = utils.RemoveFromSlice(oxygenRating, j)
				}
			}
		}
	}

	// -----------------------------------------------------------------------------
	//     - CO2 RATING -
	// -----------------------------------------------------------------------------
	co2Rating := make([]string, len(lines))
	copy(co2Rating, lines)

	for i := 0; i < len(epsilon); i += 1 {

		var ones float64
		for _, l := range co2Rating {
			if l[i] == '1' {
				ones += 1
			}
		}

		halfOfCO2 := float64(len(co2Rating)) / 2

		if len(co2Rating) == 1 {
			break
		}

		// Oxygen Generator
		for j := len(co2Rating) - 1; j >= 0; j -= 1 {
			if ones >= halfOfCO2 {
				if co2Rating[j][i] == '1' {
					co2Rating = utils.RemoveFromSlice(co2Rating, j)
				}
			} else {
				if co2Rating[j][i] == '0' {
					co2Rating = utils.RemoveFromSlice(co2Rating, j)
				}
			}
		}
	}

	oxygenDecimal, _ := strconv.ParseInt(oxygenRating[0], 2, 64)
	co2Decimal, _ := strconv.ParseInt(co2Rating[0], 2, 64)

	fmt.Printf("What is the life support rating of the submarine? ==> %d\n", oxygenDecimal*co2Decimal)
}
