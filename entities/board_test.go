package entities

import (
	"testing"
)

func TestLegalBoardSizes(t *testing.T) {
	for size := MinBoardSize; size <= MaxBoardSize; size++ {
		_, err := NewBoard(size)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestIllegalBoardSizes(t *testing.T) {
	strErr := "false negative: %d is an illegal size"

	for size := 0; size < MinBoardSize; size++ {
		_, err := NewBoard(size)
		if err == nil {
			t.Fatalf(strErr, size)
		}
	}

	for size := MaxBoardSize + 1; size < MaxBoardSize+64; size++ {
		_, err := NewBoard(size)
		if err == nil {
			t.Fatalf(strErr, size)
		}
	}
}
