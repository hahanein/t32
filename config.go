package main

import (
	"encoding/json"
	"os"
	"t32/game"
)

type Config struct {
	Size game.Size `json:"size"`

	Player1 game.Player `json:"player1"`
	Player2 game.Player `json:"player2"`
	Player3 game.Player `json:"player3"`
}

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

	return res, nil
}
