package console

import (
	"errors"
	"strconv"
	"strings"
)

// Parse parses a string in the form of NxN (where N is a natural number) into
// two integers denominating coordinates.
func Parse(s string) (int, int, error) {
	i := strings.IndexRune(s, 'x')
	if i == -1 {
		return -1, -1, errors.New("invalid input format: must be NxN where N is an integer")
	}

	sX, sY := s[:i], s[i+1:]

	x, err := strconv.Atoi(sX)
	if err != nil {
		return -1, -1, err
	}

	y, err := strconv.Atoi(sY)
	if err != nil {
		return -1, -1, err
	}

	return x, y, nil
}
