package participants

import (
	"reflect"
	"t32/game"
	"testing"
)

func TestJoin(t *testing.T) {
	r := new(spyReferee)

	p := &Participant{
		Player:  'X',
		Referee: r,
	}

	want := r.Players

	p.join()

	have := r.Players

	if reflect.DeepEqual(want, have) {
		t.Fatalf("failed to join: wanted %+v have %+v", want, have)
	}
}

func TestMove(t *testing.T) {
	r := new(spyReferee)

	coords := spyCoordinates{1, 2}

	c := &spyClient{
		Coordinates: []spyCoordinates{coords},
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

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("corrupted Move: wanted %+v have %+v", want, have)
	}
}

func TestPresent(t *testing.T) {
	want := game.Game{
		3,
		game.Players{'A', 'B', 'C'},
		game.History{game.Move{'A', 1, 2}},
	}

	r := new(spyReferee)
	c := new(spyClient)

	p := &Participant{
		Player:  'A',
		Referee: r,
		Client:  c,
	}

	p.present(want)

	have := c.Game

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("corrupted Game: wanted %+v have %+v", want, have)
	}
}
