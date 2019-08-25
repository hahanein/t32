package actors

import (
	"t32/game"
	"testing"
)

func TestIsPlayer(t *testing.T) {
	nonPlayer := &Participant{}
	player := &Participant{Player: 'X'}

	if nonPlayer.isPlayer() {
		t.Fatal("false positive: participant is not a player")
	}

	if !player.isPlayer() {
		t.Fatal("false negative: participant is a player")
	}
}

func TestJoinGame(t *testing.T) {
	r := new(SpyReferee)

	nonPlayer := &Participant{
		Referee: r,
		Client: &SpyClient{
			Players: []game.Player{game.NoPlayer},
		},
	}

	player := &Participant{
		Referee: r,
		Client: &SpyClient{
			Players: []game.Player{'X'},
		},
	}

	before := r.Players

	nonPlayer.joinGame()

	if len(before) != len(r.Players) {
		t.Fatalf("non-player MUST NOT join: wanted %v have %v", before, r.Players)
	}

	player.joinGame()

	if len(before) == len(r.Players) {
		t.Fatal("player failed to join:", r.Players)
	}
}
