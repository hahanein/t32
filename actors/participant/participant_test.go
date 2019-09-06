package participant

import (
	"t32/game"
	"testing"
)

func TestPushPlayerOnWaitingForPlayers(t *testing.T) {
	c := new(spyClient)
	r := new(spyReferee)

	r.RespStatus = game.StatusWaitingForPlayers

	_ = New('A', c, r)

	if r.Players[0] != 'A' {
		t.Fatal("participant wasn not added to players")
	}

	if !c.ReqWaitingForOthers {
		t.Fatal("must call client's WaitingForOthers method")
	}
}

func TestItsAnothersTurn(t *testing.T) {
	c := new(spyClient)
	r := new(spyReferee)

	r.RespStatus = game.StatusRunning
	r.RespWhoIsNext = 'B'

	_ = New('A', c, r)

	if !c.ReqItsAnothersTurn {
		t.Fatal("must call client's ItsAnothersTurn method")
	}
}

func TestItsYourTurn(t *testing.T) {
	c := new(spyClient)
	r := new(spyReferee)

	r.RespStatus = game.StatusRunning
	r.RespWhoIsNext = 'A'

	_ = New('A', c, r)

	if !c.ReqItsYourTurn {
		t.Fatal("must call client's ItsYourTurn method")
	}
}

func TestStalemate(t *testing.T) {
	c := new(spyClient)
	r := new(spyReferee)

	p := New('A', c, r)

	r.RespStatus = game.StatusFinish
	r.RespWinner = game.NoPlayer

	go p.Update()

	<-p.Done

	if !c.ReqStalemate {
		t.Fatal("must call client's Stalemate method")
	}
}

func TestAnotherWon(t *testing.T) {
	c := new(spyClient)
	r := new(spyReferee)

	p := New('A', c, r)

	r.RespStatus = game.StatusFinish
	r.RespWinner = 'B'

	go p.Update()

	<-p.Done

	if !c.ReqAnotherWon {
		t.Fatal("must call client's AnotherWon method")
	}
}

func TestYouWon(t *testing.T) {
	c := new(spyClient)
	r := new(spyReferee)

	p := New('A', c, r)

	r.RespStatus = game.StatusFinish
	r.RespWinner = 'A'

	go p.Update()

	<-p.Done

	if !c.ReqYouWon {
		t.Fatal("must call client's YouWon method")
	}
}

func TestFlash(t *testing.T) {
	c := new(spyClient)
	r := new(spyReferee)

	r.RespStatus = game.StatusRunning

	p := New('A', c, r)

	msg := "test flash message"

	p.Flash(game.Board{}, msg)

	if c.Message != msg {
		t.Fatalf("message delivery failed. wanted %s have %s", msg, c.Message)
	}
}
