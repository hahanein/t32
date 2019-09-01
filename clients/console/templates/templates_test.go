package templates

import (
	"strings"
	"t32/game"
	"testing"
)

func TestWhosTurn(t *testing.T) {
	var player game.Player = 'A'

	want := "It's A's turn."
	have := ItsAnothersTurn(player)

	if want != have {
		t.Fatalf("incorrect string: wanted %s have %s", want, have)
	}

	var nonPlayer game.Player = game.NoPlayer

	want = "It's nobodies turn?!"
	have = ItsAnothersTurn(nonPlayer)

	if want != have {
		t.Fatalf("incorrect string: wanted %s have %s", want, have)
	}
}

func TestBoard(t *testing.T) {
	g, err := game.Make(
		3,
		game.Players{'A', 'B', 'C'},
		game.Move{'A', 1, 2},
	)
	if err != nil {
		t.Fatal(err)
	}

	have := Board(g.Board())
	want := strings.TrimSpace(`
. . . .
. . . .
. . .A.
. . . .`)

	if want != have {
		t.Fatalf(`incorrect diagram:

wanted

%s

have

%s`, want, have)
	}
}
