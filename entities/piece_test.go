package entities

import "testing"

func TestMakeCurrentPiece(t *testing.T) {
	p, err := Game{
		Pieces: []Piece{'A', 'B', 'C'},
		History: History{
			Move{'A', 0, 0},
			Move{'B', 1, 0},
			Move{'C', 1, 1},
			Move{'A', 0, 1},
		},
	}.CurrentPiece()

	if err != nil {
		t.Fatal(err)
	}

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
