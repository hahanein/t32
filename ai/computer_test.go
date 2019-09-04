package ai

import (
	"t32/game"
	"testing"
)

func TestGenerate(t *testing.T) {
	g, _ := game.New(3)

	g.PushPlayer('A')
	g.PushPlayer('B')
	g.PushPlayer('C')
	g.PushMove(game.Move{'A', 0, 0})
	g.PushMove(game.Move{'B', 1, 0})
	g.PushMove(game.Move{'C', 0, 1})
	g.PushMove(game.Move{'A', 1, 1})
	g.PushMove(game.Move{'B', 2, 0})
	g.PushMove(game.Move{'C', 0, 2})
	g.PushMove(game.Move{'A', 2, 2})
	g.PushMove(game.Move{'B', 1, 2})

	wantedX, wantedY := 2, 1

	p := g.WhoIsNext()

	haveX, haveY := Random(g.Board(), p)

	if wantedX != haveX || wantedY != haveY {
		t.Fatalf(
			"illegal coordinates: wanted %d,%d have %d,%d",
			wantedX, wantedY,
			haveX, haveY,
		)
	}
}
