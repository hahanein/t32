package computer

import (
	"t32/game"
	"testing"
)

func TestPopCoordinates(t *testing.T) {
	c := &Computer{
		game.Game{
			3,
			game.Players{'A', 'B', 'C'},
			game.History{
				game.Move{'A', 0, 0},
				game.Move{'B', 1, 0},
				game.Move{'C', 0, 1},
				game.Move{'A', 1, 1},
				game.Move{'B', 2, 0},
				game.Move{'C', 0, 2},
				game.Move{'A', 2, 2},
				game.Move{'B', 1, 2},
			},
		},
	}

	wantedX, wantedY := 2, 1

	haveX, haveY := c.PopCoordinates()

	if wantedX != haveX || wantedY != haveY {
		t.Fatalf(
			"illegal coordinates: wanted %d,%d have %d,%d",
			wantedX, wantedY,
			haveX, haveY,
		)
	}
}
