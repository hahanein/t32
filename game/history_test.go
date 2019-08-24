package game

import "testing"

func TestIsSquareEmpty(t *testing.T) {
	g, _ := MakeGame(MinSize, Pieces{'A', 'B'})
	x, y := 0, 0

	ok := g.isSquareEmpty(x, y)
	if !ok {
		t.Fatalf("false positive: square %d,%d should be empty", x, y)
	}

	g.History = append(g.History, Move{'A', x, y})

	ok = g.isSquareEmpty(x, y)
	if ok {
		t.Fatalf("false negative: square %d,%d should be occupied", x, y)
	}
}
