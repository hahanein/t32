package referee

import (
	"t32/game"
	"testing"
)

func TestPushMove(t *testing.T) {
	s := new(Subject)
	g, _ := game.New(3)

	g.PushPlayer('X')
	g.PushPlayer('Y')
	g.PushPlayer('Z')

	r := &Referee{Subject: s, game: g}
	m := game.Move{'X', 1, 2}

	err := r.PushMove(m)
	if err != nil {
		t.Fatal(err)
	}

	err = r.PushMove(m)
	if err == nil {
		t.Fatal("pushing the same Move twice should fail")
	}
}
