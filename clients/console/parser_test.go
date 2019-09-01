package console

import "testing"

func TestParse(t *testing.T) {
	ok := "23x14"

	x, y, err := Parse(ok)
	if err != nil {
		t.Fatal(err)
	}

	if x != 23 || y != 14 {
		t.Fatalf("bad coordinates: wanted %dx%d have %dx%d", 23, 14, x, y)
	}

	bad := []string{
		"asdfg",
		"123x",
		"x123",
		"132x143x43",
		"x8xx6x5x",
	}

	for _, s := range bad {
		_, _, err := Parse(s)
		if err == nil {
			t.Fatal("invalid input must return an error:", s)
		}
	}
}
