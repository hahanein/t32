package entities

import "testing"

func TestApplyMoveToGame(t *testing.T) {
	g := &Game{
		Pieces: []Piece{'A', 'B', 'C'},
		Size:   MinSize,
	}

	m1 := Move{'A', 0, 0}

	_, err := g.Apply(m1)
	if err != nil {
		t.Fatalf("false positive: %s", err)
	}

	m2 := Move{'B', 1, 1}

	_, err = g.Apply(m2)
	if err != nil {
		t.Fatalf("false positive: %s", err)
	}

	_, err = g.Apply(m2)
	if err == nil {
		t.Fatalf("false negative: should return an error because square %d,%d is occupied", m2.X, m2.Y)
	}
}

func TestNextPiece(t *testing.T) {
	g := &Game{
		Pieces: []Piece{'A', 'B', 'C'},
		Size:   MaxSize,
	}

	p := g.NextPiece()
	if p != 'A' {
		t.Fatalf("false piece: wanted A have %c", p)
	}

	g.Apply(Move{'A', 0, 0})

	p = g.NextPiece()
	if p != 'B' {
		t.Fatalf("false piece: wanted B have %c", p)
	}
}
