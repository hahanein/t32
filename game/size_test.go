package game

import "testing"

func TestDoesSquareExist(t *testing.T) {
	var s Size = MinSize

	for x := 0; x < MinSize; x++ {
		ok := s.doesSquareExist(x, 0)
		if !ok {
			t.Fatalf("false positive: square %d,0 should exist at size %d", x, s)
		}
	}

	for x := MinSize; x < MinSize+64; x++ {
		ok := s.doesSquareExist(x, 0)
		if ok {
			t.Fatalf("false negative: square %d,0 should not exist at size %d", x, s)
		}
	}
}
