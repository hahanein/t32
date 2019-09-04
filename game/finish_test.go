package game

import (
	"reflect"
	"testing"
)

func TestStalemate(t *testing.T) {
	g := &Game{
		size:    3,
		players: Players{'A', 'B', 'C'},
		history: History{
			Move{'A', 0, 0},
			Move{'B', 0, 1},
			Move{'C', 0, 2},
			Move{'A', 1, 2},
			Move{'B', 2, 2},
			Move{'C', 2, 1},
			Move{'A', 2, 0},
			Move{'B', 1, 0},
			Move{'C', 1, 1},
		},
	}

	if !g.stalemate() {
		t.Fatal("must be true since no more Moves are possible")
	}
}

func TestWinnerRow(t *testing.T) {
	wonRow := Row{'X', 'X', 'X', 'X'}

	p := wonRow.winner()
	if p != 'X' {
		t.Fatal("should return X since Row is occupied by that player only")
	}

	lostRows := []Row{
		Row{NoPlayer, NoPlayer, NoPlayer},
		Row{'A', 'B', 'C'},
	}

	for _, r := range lostRows {
		p := r.winner()
		if p != NoPlayer {
			t.Fatal("should not return a winner. returned:", p)
		}
	}
}

func TestBoardRotate(t *testing.T) {
	have := Board{
		[]Player{1, 2, 3},
		[]Player{8, 9, 4},
		[]Player{7, 6, 5},
	}.rotate()

	want := Board{
		[]Player{3, 4, 5},
		[]Player{2, 9, 6},
		[]Player{1, 8, 7},
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("reversing board failed. wanted %+v have %+v", want, have)
	}

	have = Board{
		[]Player{1, 0, 0},
		[]Player{0, 1, 0},
		[]Player{0, 0, 1},
	}.rotate()

	want = Board{
		[]Player{0, 0, 1},
		[]Player{0, 1, 0},
		[]Player{1, 0, 0},
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("reversing board failed. wanted %+v have %+v", want, have)
	}
}

func TestDiagonal(t *testing.T) {
	have := Board{
		Row{1, 0, 0},
		Row{0, 1, 0},
		Row{0, 0, 1},
	}.diagonal()

	want := Row{1, 1, 1}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("finding diagonal failed. wanted %+v have %+v", want, have)
	}
}

func TestWinner(t *testing.T) {
	won := Game{
		size:    3,
		players: Players{'A', 'B', 'C'},
		history: History{
			Move{'A', 0, 0},
			Move{'B', 1, 0},
			Move{'C', 2, 0},
			Move{'A', 0, 1},
			Move{'B', 1, 1},
			Move{'C', 2, 1},
			Move{'A', 0, 2},
		},
	}

	if won.Status() != StatusFinish {
		t.Fatal("incorrect game state. A won.")
	}

	if p := won.Winner(); p != 'A' {
		t.Fatalf("incorrect winner. wanted A have %c", p)
	}
}
