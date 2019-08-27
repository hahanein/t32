package game

import (
	"reflect"
	"testing"
)

var gmsLegal = []Game{
	Game{
		3,
		Players{'A', 'B', 'C'},
		History{},
	},
	Game{
		10,
		Players{'Ä', 'B', '😛'},
		History{
			Move{'Ä', 0, 0},
		},
	},
	Game{
		5,
		Players{'1', '2', '3'},
		History{
			Move{'1', 0, 0},
			Move{'2', 2, 0},
			Move{'3', 0, 1},
			Move{'1', 0, 2},
			Move{'2', 3, 0},
			Move{'3', 0, 3},
		},
	},
}

var gmsIllegal = []Game{
	Game{
		1,
		Players{'A', 'B', 'C'},
		History{},
	},
	Game{
		5,
		Players{'A', 'B'},
		History{},
	},
	Game{
		10,
		Players{'Ä', 'B', NoPlayer},
		History{
			Move{'Ä', 0, 0},
		},
	},
	Game{
		5,
		Players{'1', '2', '3'},
		History{
			Move{'1', 0, 0},
			Move{'2', 2, 0},
			Move{'3', 0, 1},
			Move{'1', 0, 2},
			Move{'2', 3, 0},
			Move{'3', 0, 3},
			Move{'1', 0, 0},
		},
	},
	Game{
		4,
		Players{'1', '2', '3'},
		History{
			Move{'1', 0, 0},
			Move{'2', 2, 0},
			Move{'3', 0, 1},
			Move{'2', 3, 0},
			Move{'1', 0, 2},
			Move{'3', 0, 3},
			Move{'1', 0, 0},
		},
	},
}

func TestLegalSizes(t *testing.T) {
	var s Size

	for s = MinSize; s <= MaxSize; s++ {
		_, err := New(s)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestIllegalSizes(t *testing.T) {
	var s Size

	strErr := "false negative: %d is an illegal size"

	for s = 0; s < MinSize; s++ {
		_, err := New(s)
		if err == nil {
			t.Fatalf(strErr, s)
		}
	}

	for s = MaxSize + 1; s < MaxSize+64; s++ {
		_, err := New(s)
		if err == nil {
			t.Fatalf(strErr, s)
		}
	}
}

func TestMake(t *testing.T) {
	left := &Game{size: 3}

	right, err := New(3)
	if err != nil {
		t.Fatal("false positive:", err)
	}

	if !reflect.DeepEqual(left, right) {
		t.Fatalf("corrupted Game: wanted size %d have size %d", left, right)
	}
}
