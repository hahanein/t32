package game

import (
	"encoding/json"
	"errors"
)

// UnmarshalJSON parses a JSON-encoded Size and stores the result in the value
// pointed to by s.
func (s *Size) UnmarshalJSON(bytes []byte) error {
	var raw int

	err := json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}

	*s = Size(raw)

	return nil
}

// UnmarshalJSON parses a JSON-encoded Player and stores the result in the
// value pointed to by p.
func (p *Player) UnmarshalJSON(bytes []byte) error {
	var raw string

	err := json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}

	if len(raw) == 1 {
		*p = Player(raw[0])

		return nil
	}

	return errors.New("invalid length")
}
