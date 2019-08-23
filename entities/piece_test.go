package entities

import "testing"

func TestMakeCurrentPiece(t *testing.T) {
	p, err := MakeCurrentPiece(
		[]Piece{'A', 'B', 'C'},
		Move{'A', 0, 0},
		Move{'B', 1, 0},
		Move{'C', 1, 1},
		Move{'A', 0, 1},
	)

	if err != nil {
		t.Fatal(err)
	}

	if p != 'B' {
		t.Fatalf("false piece: wanted B have %c", p)
	}

}
