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

func TestIsSquareEmpty(t *testing.T) {
	b, _ := NewBoard(MinBoardSize)
	x, y := 0, 0

	ok := b.isSquareEmpty(x, y)
	if !ok {
		t.Fatalf("false positive: square %d,%d should be empty", x, y)
	}

	(*b)[x][y] = 'X'

	ok = b.isSquareEmpty(x, y)
	if ok {
		t.Fatalf("false negative: square %d,%d should be occupied", x, y)
	}
}

func TestDoesSquareExist(t *testing.T) {
	b, _ := NewBoard(MinBoardSize)

	for x := 0; x < MinBoardSize; x++ {
		ok := b.doesSquareExist(x, 0)
		if !ok {
			t.Fatalf("false positive: square %d,0 should exist", x)
		}
	}

	for x := MinBoardSize; x < MinBoardSize+64; x++ {
		ok := b.doesSquareExist(x, 0)
		if ok {
			t.Fatalf("false negative: square %d,0 should not exist", x)
		}
	}
}

func TestApplyMove(t *testing.T) {
	b, _ := NewBoard(MinBoardSize)
	m := Move{'X', 0, 0}

	err := b.Apply(m)
	if err != nil {
		t.Fatalf("false positive: %s", err)
	}

	err = b.Apply(m)
	if err == nil {
		t.Fatalf("false negative: should return an error because square %d,%d is occupied", m.X, m.Y)
	}
}
