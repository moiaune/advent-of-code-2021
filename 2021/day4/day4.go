package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madsaune/advent-of-code/internal/utils"
)

// square represents each number on a bingo board
type square struct {
	value  int
	marked bool
}

// board represents the bingo board itself
type board struct {
	grid []square
}

func main() {
	var input string

	flag.StringVar(&input, "input", "input.txt", "Path to input file")
	flag.Parse()

	lines, err := utils.Readlines(input)
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	// -----------------------------------------------------------------------------
	//     - PARSE MOVES -
	// -----------------------------------------------------------------------------
	moves := strings.Split(lines[0], ",")

	// -----------------------------------------------------------------------------
	//     - PARSE BOARDS -
	// -----------------------------------------------------------------------------
	var boards []board
	for i := 2; i <= len(lines); i += 6 {
		var b board

		for j := 0; j < 6; j += 1 {
			if i+j == len(lines) {
				break
			}

			rowStr := strings.Split(lines[i+j], " ")
			for _, v := range rowStr {
				if v == " " || v == "" {
					continue
				}

				tmp, _ := strconv.Atoi(v)
				n := square{
					value: tmp,
				}

				b.grid = append(b.grid, n)
			}
		}

		boards = append(boards, b)
	}

	// -----------------------------------------------------------------------------
	//     - PLAY ROUNDS -
	// -----------------------------------------------------------------------------

	var hasWinner bool
	var lastMove int
	var winningBoard int

	for _, n := range moves {
		if hasWinner {
			break
		}

		// convert move to int
		asInt, _ := strconv.Atoi(n)

		// --- MARK NUMBERS ON BOARDS
		for bi := range boards {
			for gi := range boards[bi].grid {
				if boards[bi].grid[gi].value == asInt {
					boards[bi].grid[gi].marked = true
				}
			}
		}

		// --- CHECK FOR WINNERS
		for bi := range boards {
			grid := boards[bi].grid

			if isRowWinner(grid) || isColWinner(grid) {
				hasWinner = true
				winningBoard = bi
				break
			}
		}

		lastMove = asInt
	}

	sum := sumOfUnmarkedSquares(boards[winningBoard].grid)

	fmt.Println("Unmarked sum: ", sum)
	fmt.Println("Lastmove: ", lastMove)

	fmt.Println("What will your final score be if you choose that board? ==>", lastMove*sum)
	fmt.Println("----------------------")

	// -----------------------------------------------------------------------------
	//     - PART 2 -
	//
	// Figure out which board will win last.
	// -----------------------------------------------------------------------------

	boards2 := make([]board, len(boards))
	copy(boards2, boards)
	var lastBoardToWin board
	var lastMoveToWin int

	for _, n := range moves {

		// convert move to int
		asInt, _ := strconv.Atoi(n)

		if len(boards2) == 0 {
			break
		}

		// --- MARK NUMBERS ON BOARDS
		for bi := range boards2 {
			for gi := range boards2[bi].grid {
				if boards2[bi].grid[gi].value == asInt {
					boards2[bi].grid[gi].marked = true
				}
			}
		}

		// --- CHECK FOR WINNERS
		for i := len(boards2) - 1; i >= 0; i -= 1 {
			grid := boards2[i].grid

			if isRowWinner(grid) || isColWinner(grid) {
				lastBoardToWin = boards2[i]
				lastMoveToWin = asInt
				boards2 = removeBoard(boards2, i)
			}
		}

		lastMove = asInt
	}

	sumUnmarked := sumOfUnmarkedSquares(lastBoardToWin.grid)

	fmt.Println("Unmarked sum: ", sumUnmarked)
	fmt.Println("Lastmove: ", lastMoveToWin)
	fmt.Println("What will your final score be if you choose that board? ==>", lastMoveToWin*sumUnmarked)
}

func removeBoard(b []board, idx int) []board {
	copy(b[idx:], b[idx+1:]) // Shift a[i+1:] left one index.
	b[len(b)-1] = board{}    // Erase last element (write zero value).
	b = b[:len(b)-1]         // Truncate slice.

	return b
}

func sumOfUnmarkedSquares(g []square) int {
	var sum int
	for _, n := range g {
		if n.marked == false {
			sum += n.value
		}
	}

	return sum
}

func isRowWinner(row []square) bool {
	for i := 0; i < len(row); i += 5 {
		end := i + 5
		start := end - 5

		isWinningRow := true
		row := row[start:end]
		for _, n := range row {
			if !n.marked {
				// if atleast one number in row is false
				// then its not a winning row
				isWinningRow = false
				break
			}
		}

		if isWinningRow {
			return true
		}
	}

	return false
}

func isColWinner(g []square) bool {
	for col := 0; col < 5; col += 1 {

		var rows []square
		isWinningCol := true

		rows = append(rows, g[col])
		if !g[col].marked {
			isWinningCol = false
		}

		for i := 5; i < len(g); i += 5 {
			rows = append(rows, g[col+i])
			if !g[col+i].marked {
				isWinningCol = false
			}
		}

		if isWinningCol {
			return true
		}

	}

	return false
}

func getGridRow(b board, row int) []square {
	end := row * 5
	start := end - 5

	return b.grid[start:end]
}
