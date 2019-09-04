package game

import (
	"encoding/json"
	"errors"
)

func (s *Size) UnmarshalJSON(bytes []byte) error {
	var raw int

	err := json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}

	*s = Size(raw)

	return nil
}

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
