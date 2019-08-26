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
