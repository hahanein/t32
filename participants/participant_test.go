package actors

import (
	"t32/game"
	"testing"
)

func TestJoin(t *testing.T) {
	r := new(SpyReferee)

	p := &Participant{
		Player:  'X',
		Referee: r,
	}

	before := r.Players

	p.join()

	if len(before) == len(r.Players) {
		t.Fatal("failed to join:", r.Players)
	}
}

func TestMove(t *testing.T) {
	r := new(SpyReferee)

	coords := SpyCoordinates{1, 2}

	c := &SpyClient{
		Coordinates: []SpyCoordinates{coords},
	}

	p := &Participant{
		Player:  'X',
		Referee: r,
		Client:  c,
	}

	p.move()

	if len(r.History) == 0 {
		t.Fatal("failed to send next Move")
	}

	have := r.History[0]
	want := game.Move{p.Player, coords.X, coords.Y}

	if want.Player != have.Player || want.X != have.X || want.Y != have.Y {
		t.Fatalf("corrupted Move: wanted %+v have %+v", want, have)
	}
}
