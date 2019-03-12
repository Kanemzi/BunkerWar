package main

import (
	"errors"
)

// PlayerBase represents a base created by a player in the game world
type PlayerBase struct {
	// Physical coordinates of the base
	Loc Location

	// Logins of the bases's control panel
	Login, Password string

	// The player who constructed the base
	Owner *Player

	// Network connections to other bases
	NetworkLinks map[int]NetworkLink

	// Technical level of the base
	TechLevel float64
}

// NewPlayerBase creates and returns a pointer on a new base
func NewPlayerBase(loc Location, login, password string) *PlayerBase {
	base := new(PlayerBase)
	base.Loc = loc
	base.Login = login
	base.Password = password
	return base
}

// AssignPlayer defines what player has the ownership on a base
// a base can be destroyed but can't be taken to another player
// if the base already has an owner, an error will be returned
func (base *PlayerBase) AssignPlayer(player *Player) error {
	if base.Owner != nil {
		return errors.New("This base is already assigned to the player " + outPlayerName(base.Owner.Name))
	}

	base.Owner = player
	return nil
}
