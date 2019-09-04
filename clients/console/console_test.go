package console

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"t32/game"
	"testing"
)

func TestDumbHandlers(t *testing.T) {
	want := "OK"

	w := new(bytes.Buffer)
	c := &Console{sync.Mutex{}, &spyTemplates{want}, w, nil}

	epyB := game.Board{}
	epyP := game.NoPlayer

	c.WaitingForOthers()
	have := w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	c.ItsAnothersTurn(epyB, epyP)
	have = w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	c.Stalemate(epyB)
	have = w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	c.AnotherWon(epyB, epyP)
	have = w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	c.YouWon(epyB, epyP)
	have = w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()
}

func TestItsYourTurn(t *testing.T) {
	epyB := game.Board{}
	epyP := game.NoPlayer

	want := "OK"
	wantX, wantY := 1, 2

	w := new(bytes.Buffer)
	r := strings.NewReader(fmt.Sprintf("%dx%d", wantX, wantY))
	c := &Console{sync.Mutex{}, &spyTemplates{want}, w, r}

	haveX, haveY := c.ItsYourTurn(epyB, epyP)

	have := w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	if haveX != wantX || haveY != wantY {
		t.Fatalf("handler is broken: wanted %dx%d have %dx%d", wantX, wantY, haveX, haveX)
	}
}
