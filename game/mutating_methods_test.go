package game

import "testing"

func TestPushPlayer(t *testing.T) {
	var ps Players

	ps2, err := ps.PushPlayer('X')
	if err != nil {
		t.Fatal(err)
	}

	if len(ps2) == 0 || ps2[0] != 'X' {
		t.Fatal("failed to add Player to Players list")
	}
}

func TestPushMove(t *testing.T) {
	g := Game{3, Players{'A', 'B', 'C'}, History{}}
	m := Move{'A', 1, 2}

	err := g.PushMove(m)
	if err != nil {
		t.Fatal(err)
	}
}
