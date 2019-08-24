package game

import "testing"

func TestCurrentPlayer(t *testing.T) {
	g, _ := Make(MinSize, Players{'A', 'B', 'C'},
		Move{'A', 0, 0},
		Move{'B', 1, 0},
		Move{'C', 1, 1},
		Move{'A', 0, 1},
	)

	p := g.CurrentPlayer()

	if p != 'B' {
		t.Fatalf("false player: wanted B have %c", p)
	}
}

func TestHasUniqItemsOnly(t *testing.T) {
	ps := Players{'A', 'B', 'C'}

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
