package game

import "testing"

func TestCurrentPiece(t *testing.T) {
	g, _ := Make(MinSize, Pieces{'A', 'B', 'C'},
		Move{'A', 0, 0},
		Move{'B', 1, 0},
		Move{'C', 1, 1},
		Move{'A', 0, 1},
	)

	p := g.CurrentPiece()

	if p != 'B' {
		t.Fatalf("false piece: wanted B have %c", p)
	}
}

func TestHasUniqItemsOnly(t *testing.T) {
	ps := Pieces{'A', 'B', 'C'}

	ok := ps.hasUniqItemsOnly()
	if !ok {
		t.Fatal("false positive: list has no duplicate items")
	}

	ps = append(ps, 'C')

	ok = ps.hasUniqItemsOnly()
	if ok {
		t.Fatal("false negative: list has duplicate items")
	}
}
