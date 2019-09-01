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

	for _, r := range raw {
		*p = Player(r)

		return nil
	}

	return errors.New("invalid length")
}
