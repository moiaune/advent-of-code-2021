package utils

import (
	"bufio"
	"log"
	"os"
)

func LoadFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("ERR: Could not open input file: %v", err)
	}

	return f
}

func Readlines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func BinaryToDecimal(binary string) int64 {
	var result int64
	var bit int

	n := len(binary) - 1
	for n >= 0 {
		if binary[n] == '1' {
			result += (1 << (bit))
		}

		n -= 1
		bit += 1
	}

	return result
}

func RemoveFromSlice(slice []string, index int) []string {
	copy(slice[index:], slice[index+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = ""             // Erase last element (write zero value).
	slice = slice[:len(slice)-1]         // Truncate slice.

	return slice
}
