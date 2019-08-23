package entities

import (
	"testing"
)

func TestLegalBoardSizes(t *testing.T) {
	for size := MinSize; size <= MaxSize; size++ {
		_, err := MakeBoard(size)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestIllegalBoardSizes(t *testing.T) {
	strErr := "false negative: %d is an illegal size"

	for size := 0; size < MinSize; size++ {
		_, err := MakeBoard(size)
		if err == nil {
			t.Fatalf(strErr, size)
		}
	}

	for size := MaxSize + 1; size < MaxSize+64; size++ {
		_, err := MakeBoard(size)
		if err == nil {
			t.Fatalf(strErr, size)
		}
	}
}

func TestIsSquareEmpty(t *testing.T) {
	b, _ := MakeBoard(MinSize)
	x, y := 0, 0

	ok := b.isSquareEmpty(x, y)
	if !ok {
		t.Fatalf("false positive: square %d,%d should be empty", x, y)
	}

	b[x][y] = 'X'

	ok = b.isSquareEmpty(x, y)
	if ok {
		t.Fatalf("false negative: square %d,%d should be occupied", x, y)
	}
}

func TestDoesSquareExist(t *testing.T) {
	b, _ := MakeBoard(MinSize)

	for x := 0; x < MinSize; x++ {
		ok := b.doesSquareExist(x, 0)
		if !ok {
			t.Fatalf("false positive: square %d,0 should exist", x)
		}
	}

	for x := MinSize; x < MinSize+64; x++ {
		ok := b.doesSquareExist(x, 0)
		if ok {
			t.Fatalf("false negative: square %d,0 should not exist", x)
		}
	}
}

func TestApplyMoveToBoard(t *testing.T) {
	b, _ := MakeBoard(MinSize)

	_, err := b.apply()
	if err != nil {
		t.Fatalf("false positive: %s", err)
	}

	m := Move{'X', 0, 0}

	b2, err := b.apply(m)
	if err != nil {
		t.Fatalf("false positive: %s", err)
	}

	_, err = b2.apply(m)
	if err == nil {
		t.Fatalf("false negative: should return an error because square %d,%d is occupied", m.X, m.Y)
	}
}
