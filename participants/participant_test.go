package actors

import (
	"testing"
)

func TestJoinGame(t *testing.T) {
	r := new(SpyReferee)

	p := &Participant{
		Player:  'X',
		Referee: r,
	}

	before := r.Players

	p.joinGame()

	if len(before) == len(r.Players) {
		t.Fatal("failed to join:", r.Players)
	}
}
