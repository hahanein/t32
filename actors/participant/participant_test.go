package participant

import "testing"

func TestPushPlayerOnWaitingForPlayers(t *testing.T) {
	c := new(spyClient)
	r := new(spyReferee)
	_ = New('A', c, r)

	if r.Players[0] != 'A' {
		t.Fatal("participant wasn not added to players")
	}
}
