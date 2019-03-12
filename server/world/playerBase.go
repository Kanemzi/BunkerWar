package world

import (
	"errors"

	"github.com/SimonROZEC/bunker/server/globals"
	"github.com/SimonROZEC/bunker/server/util"
)

// PlayerBase represents a base created by a player in the game world
type PlayerBase struct {
	// Physical coordinates of the base
	Loc util.Location

	// Logins of the bases's control panel
	Login, Password string

	// The player who constructed the base
	Owner *Player

	// Network connections to other bases
	NetworkLinks map[int]Link

	// Technical level of the base
	TechLevel float64
}

// NewPlayerBase creates and returns a pointer on a new base
func NewPlayerBase(loc util.Location, login, password string) *PlayerBase {
	base := new(PlayerBase)
	base.Loc = loc
	base.Login = login
	base.Password = password
	base.NetworkLinks = make(map[int]Link, 0)
	base.TechLevel = 0
	return base
}

// AssignPlayer defines what player has the ownership on a base
// a base can be destroyed but can't be taken to another player
// if the base already has an owner, an error will be returned
func (base *PlayerBase) AssignPlayer(player *Player) error {
	if base.Owner != nil {
		return errors.New("This base is already assigned to the player " + util.OutPlayerName(base.Owner.Name))
	}

	base.Owner = player
	return nil
}

// AddLink adds a new link in the list of interfaces of the base
func (base *PlayerBase) AddLink(link *Link) {
	base.NetworkLinks = 
}

// Connect creates a connection between two bases
// returns the created Link or an error if the link already exists
func (base *PlayerBase) Connect(target *PlayerBase) (*Link, error) {
	var link *Link
	exists, _ := LinkExists(make([]*PlayerBase, 5), base, target)
	if exists {
		return link, errors.New("The base " + util.OutBaseName(base.Login) + "is already connected to " + util.OutBaseName(target.Login))
	}

	link = NewLink(base, target, globals.InitialNetworkSpeed)

	return link, nil

}
