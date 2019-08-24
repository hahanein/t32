package game

import (
	"testing"
)

func TestLegalSizes(t *testing.T) {
	var s Size

	for s = MinSize; s <= MaxSize; s++ {
		_, err := Make(s, Pieces{'A', 'B', 'C'})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestIllegalSizes(t *testing.T) {
	var s Size

	strErr := "false negative: %d is an illegal size"

	for s = 0; s < MinSize; s++ {
		_, err := Make(s, Pieces{'A', 'B', 'C'})
		if err == nil {
			t.Fatalf(strErr, s)
		}
	}

	for s = MaxSize + 1; s < MaxSize+64; s++ {
		_, err := Make(s, Pieces{'A', 'B', 'C'})
		if err == nil {
			t.Fatalf(strErr, s)
		}
	}
}

func TestValidate(t *testing.T) {
	err := Validate(MinSize, Pieces{'A', 'B', 'C'})
	if err != nil {
		t.Fatal("false positive:", err)
	}
}

func TestMake(t *testing.T) {
	left := Game{3, Pieces{'A', 'B', 'C'}, History{}}

	right, err := Make(3, Pieces{'A', 'B', 'C'})
	if err != nil {
		t.Fatal("false positive:", err)
	}

	if left.Size != right.Size {
		t.Fatalf("corrupted Game: wanted size %d have size %d", left.Size, right.Size)
	}

	lenLeft, lenRight := len(left.Pieces), len(right.Pieces)

	if lenLeft != lenRight {
		t.Fatalf("corrupted Game: wanted %d number of pieces have %d", lenLeft, lenRight)
	}
}
