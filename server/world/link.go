package world

// InterfaceStatus represents the status of the interface of a base in a network connection
type InterfaceStatus int

const (
	// CLOSED means that the interface is closed, data can't circulate between the two bases
	CLOSED InterfaceStatus = 0

	// OPENED means that the interface is opened, if the two interfaces of a connecion are opened, the two bases can share information
	OPENED InterfaceStatus = 1
)

// Link represents a network connection between two bases
type Link struct {
	OwnerBase, TargetBase     *PlayerBase
	OwnerStatus, TargetStatus InterfaceStatus
	Speed                     float64
}

// NewLink creates and returns a pointer on a new Link
func NewLink(owner, target *PlayerBase, speed float64) *Link {
	link := new(Link)
	link.OwnerBase = owner
	link.TargetBase = target
	link.Speed = speed
	link.OwnerStatus = OPENED
	link.TargetStatus = OPENED
	return link
}

// LinkExists returns true if a link between two bases has already been created
// if the link exists, it's also returned
func LinkExists(bases []*PlayerBase, a *PlayerBase, b *PlayerBase) (bool, Link) {
	var nillink Link
	for _, b := range bases {
		for _, l := range b.NetworkLinks {
			owa := l.OwnerBase == a && l.TargetBase == b
			owb := l.OwnerBase == b && l.TargetBase == a
			if owa || owb {
				return true, l
			}
		}
	}
	return false, nillink
}
