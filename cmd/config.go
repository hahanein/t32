package main

import (
	"encoding/json"
	"os"
	"t32/game"
)

// Config is used to configure the Board Size as well as the Player characters.
type Config struct {
	Size game.Size `json:"size"`

	Player1 game.Player `json:"player1"`
	Player2 game.Player `json:"player2"`
	Player3 game.Player `json:"player3"`
}

// configFromFilepath parses a config file in a given path and on success
// returns a Config struct. Otherwise it returns an error.
func configFromFilepath(fp string) (Config, error) {
	var res Config

	f, err := os.Open(fp)
	if err != nil {
		return res, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&res)
	if err != nil {
		return res, err
	}

	if !res.Player1.IsLegal() {
		return res, game.ErrPlayerIllegal
	}

	if !res.Player2.IsLegal() {
		return res, game.ErrPlayerIllegal
	}

	if !res.Player3.IsLegal() {
		return res, game.ErrPlayerIllegal
	}

	// Check if all Players are unique and return an error if they are not.
	m := make(map[game.Player]struct{})
	m[res.Player1] = struct{}{}
	m[res.Player2] = struct{}{}
	m[res.Player3] = struct{}{}
	if len(m) != 3 {
		return res, game.ErrPlayerExists
	}

	return res, nil
}
