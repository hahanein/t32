package game

import (
	"testing"
)

// TODO: Function which checks if a Game has finished. A winner should be
// determined as well.

func TestLegalSizes(t *testing.T) {
	for size := MinSize; size <= MaxSize; size++ {
		_, err := MakeGame(size, Pieces{'A', 'B', 'C'})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestIllegalSizes(t *testing.T) {
	strErr := "false negative: %d is an illegal size"

	for size := 0; size < MinSize; size++ {
		_, err := MakeGame(size, Pieces{'A', 'B', 'C'})
		if err == nil {
			t.Fatalf(strErr, size)
		}
	}

	for size := MaxSize + 1; size < MaxSize+64; size++ {
		_, err := MakeGame(size, Pieces{'A', 'B', 'C'})
		if err == nil {
			t.Fatalf(strErr, size)
		}
	}
}

func TestDoesSquareExist(t *testing.T) {
	g, _ := MakeGame(MinSize, Pieces{'A', 'B', 'C'})

	for x := 0; x < MinSize; x++ {
		ok := g.doesSquareExist(x, 0)
		if !ok {
			t.Fatalf("false positive: square %d,0 should exist", x)
		}
	}

	for x := MinSize; x < MinSize+64; x++ {
		ok := g.doesSquareExist(x, 0)
		if ok {
			t.Fatalf("false negative: square %d,0 should not exist", x)
		}
	}
}

func TestApply(t *testing.T) {
	g, _ := MakeGame(MinSize, Pieces{'A', 'B', 'C'})

	_, err := g.apply()
	if err != nil {
		t.Fatalf("false positive: %s", err)
	}

	m := Move{'A', 0, 0}

	g2, err := g.apply(m)
	if err != nil {
		t.Fatalf("false positive: %s", err)
	}

	_, err = g2.apply(m)
	if err == nil {
		t.Fatalf("false negative: should return an error because square %d,%d is occupied", m.X, m.Y)
	}
}

func TestMakeGame(t *testing.T) {
	left := Game{3, Pieces{'A', 'B', 'C'}, History{}}

	right, err := MakeGame(3, Pieces{'A', 'B', 'C'})
	if err != nil {
		t.Fatal(err)
	}

	if left.Size != right.Size {
		t.Fatalf("corrupted Game: wanted size %d have size %d", left.Size, right.Size)
	}

	lenLeft, lenRight := len(left.Pieces), len(right.Pieces)

	if lenLeft != lenRight {
		t.Fatalf("corrupted Game: wanted %d number of pieces have %d", lenLeft, lenRight)
	}
}
