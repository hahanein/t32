package actors

import "t32/game"

type SpyReferee struct {
	game.Game
	Err error
}

func (r *SpyReferee) PushMove(m game.Move) error {
	r.History = append(r.History, m)
	return nil
}

func (r *SpyReferee) GetGame() (game.Game, error) {
	return r.Game, nil
}

func (r *SpyReferee) SetPlayer(p game.Player) error {
	r.Players = append(r.Players, p)
	return nil
}

func (r *SpyReferee) Fatal(err error) {
	r.Err = err
}
