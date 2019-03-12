package world

import "github.com/SimonROZEC/bunker/server/util"

// Player is a player in the game
type Player struct {
	// Physical coordinates of the player
	Loc util.Location

	// The base where the player is located
	Base *PlayerBase

	// The name of the player
	Name string
}

// NewPlayer creates and returns a pointer on a new Player
func NewPlayer(name string, loc util.Location, base *PlayerBase) *Player {
	player := new(Player)
	player.Name = name
	player.Loc = loc
	player.Base = base
	return player
}
