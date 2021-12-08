package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/madsaune/advent-of-code/internal/utils"
)

type point struct {
	x     int
	y     int
	count int
}

func main() {
	var input string
	var gridSize int

	flag.StringVar(&input, "input", "input.txt", "Path to input file")
	flag.IntVar(&gridSize, "size", 1000, "Size of grid")
	flag.Parse()

	lines, err := utils.Readlines(input)
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	// -----------------------------------------------------------------------------
	//     - CREATE GRID -
	// -----------------------------------------------------------------------------
	var grid [][]point
	for i := 0; i < gridSize; i += 1 {
		var row []point
		for j := 0; j < gridSize; j += 1 {
			row = append(row, point{
				x: j,
				y: i,
			})
		}

		grid = append(grid, row)
	}

	// -----------------------------------------------------------------------------
	//     - PLOT LINES -
	// -----------------------------------------------------------------------------
	for _, l := range lines {
		parsed := strings.Split(l, " -> ")
		fromParsed := strings.Split(parsed[0], ",")
		toParsed := strings.Split(parsed[1], ",")

		fromX, _ := strconv.Atoi(fromParsed[0])
		fromY, _ := strconv.Atoi(fromParsed[1])

		toX, _ := strconv.Atoi(toParsed[0])
		toY, _ := strconv.Atoi(toParsed[1])

		if fromX == toX {
			diff := math.Abs(float64(fromY) - float64(toY))
			// fmt.Printf("fy: %d, ty: %d, diff: %d\n", fromY, toY, int(diff))

			// I think we can do something smarter with the diff here..
			if fromY < toY {
				for i := 0; i <= int(diff); i += 1 {
					grid[fromX][fromY+i].count += 1
				}
			} else {
				for i := 0; i <= int(diff); i += 1 {
					grid[fromX][fromY-i].count += 1
				}
			}
		} else if fromY == toY {
			diff := math.Abs(float64(fromX) - float64(toX))
			// fmt.Printf("fx: %d, tx: %d, diff: %d\n", fromX, toX, int(diff))

			// I think we can do something smarter with the diff here..
			if fromX < toX {
				for i := 0; i <= int(diff); i += 1 {
					grid[fromX+i][fromY].count += 1
				}
			} else {
				for i := 0; i <= int(diff); i += 1 {
					grid[fromX-i][fromY].count += 1
				}
			}
		} else {
			// -----------------------------------------------------------------------------
			//     - PART 2 -
			//
			// ITS DIAGONAL
			// -----------------------------------------------------------------------------
			diff := math.Abs(float64(fromX) - float64(toX))
			// fmt.Printf("f: %d,%d, t: %d,%d, diffX: %d, diffY: %d\n", fromX, fromY, toX, toY, int(diffX), int(diffY))

			if (fromX > toX) && (fromY < toY) {
				for i := 0; i <= int(diff); i += 1 {
					grid[fromX-i][fromY+i].count += 1
				}
			} else if (fromX < toX) && (fromY > toY) {
				for i := 0; i <= int(diff); i += 1 {
					grid[fromX+i][fromY-i].count += 1
				}
			} else if (fromX < toX) && (fromY < toY) {
				for i := 0; i <= int(diff); i += 1 {
					grid[fromX+i][fromY+i].count += 1
				}
			} else {
				for i := 0; i <= int(diff); i += 1 {
					grid[fromX-i][fromY-i].count += 1
				}
			}
		}
	}

	// -----------------------------------------------------------------------------
	//     - CALCULATE OVERLAPPING LINES -
	// -----------------------------------------------------------------------------
	var linesOverlapping int
	for i := 0; i < gridSize; i += 1 {
		for j := 0; j < gridSize; j += 1 {
			if grid[i][j].count > 1 {
				linesOverlapping += 1
			}
		}
	}

	fmt.Println("Lines overlapping:", linesOverlapping)
}
