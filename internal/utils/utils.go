package utils

import (
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
