package game

import "testing"

func TestWhosNext(t *testing.T) {
	g := &Game{
		MinSize,
		Players{'A', 'B', 'C'},
		History{
			Move{'A', 0, 0},
			Move{'B', 1, 0},
			Move{'C', 1, 1},
			Move{'A', 0, 1},
		},
	}

	if p := g.WhoIsNext(); p != 'B' {
		t.Fatalf("false player: wanted B have %c", p)
	}
}
