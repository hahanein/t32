package console

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"t32/game"
	"testing"
)

func TestDumbHandlers(t *testing.T) {
	want := "OK"

	w := new(bytes.Buffer)
	c := &Console{&spyTemplates{want}, w, nil}

	epyB := game.Board{}
	epyP := game.NoPlayer

	c.WaitingForOthers(context.Background())
	have := w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	c.ItsAnothersTurn(context.Background(), epyB, epyP)
	have = w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	c.Stalemate(context.Background(), epyB)
	have = w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	c.AnotherWon(context.Background(), epyB, epyP)
	have = w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	c.YouWon(context.Background(), epyB, epyP)
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
	c := &Console{&spyTemplates{want}, w, r}

	haveX, haveY := c.ItsYourTurn(context.Background(), epyB, epyP)

	have := w.String()
	if have != want {
		t.Fatalf("handler is broken: wanted %s have %s", want, have)
	}
	w.Reset()

	if haveX != wantX || haveY != wantY {
		t.Fatalf("handler is broken: wanted %dx%d have %dx%d", wantX, wantY, haveX, haveX)
	}
}
