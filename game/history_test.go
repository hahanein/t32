package game

import "testing"

func TestIsSquareEmpty(t *testing.T) {
	h := History{}
	x, y := 0, 0

	ok := h.isSquareEmpty(x, y)
	if !ok {
		t.Fatalf("false positive: square %d,%d should be empty", x, y)
	}

	h = append(h, Move{'A', x, y})

	ok = h.isSquareEmpty(x, y)
	if ok {
		t.Fatalf("false negative: square %d,%d should be occupied", x, y)
	}
}
