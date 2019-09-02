package templates

import (
	"strings"
	"t32/game"
	"testing"
)

func TestBoard(t *testing.T) {
	g, _ := game.New(3)
	g.PushPlayer('A')
	g.PushPlayer('B')
	g.PushPlayer('C')
	g.PushMove(game.Move{'A', 1, 2})

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
