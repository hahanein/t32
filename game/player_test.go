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

	p, err := g.WhoIsNext()
	if err != nil {
		t.Fatal(err)
	}

	if p != 'B' {
		t.Fatalf("false player: wanted B have %c", p)
	}
}
