package main

// InterfaceStatus represents the status of the interface of a base in a network connection
type InterfaceStatus int

const (
	// CLOSED means that the interface is closed, data can't circulate between the two bases
	CLOSED InterfaceStatus = 0

	// OPENED means that the interface is opened, if the two interfaces of a connecion are opened, the two bases can share information
	OPENED InterfaceStatus = 1
)

// NetworkLink represents a network connection between two bases
type NetworkLink struct {
	OwnerBase, TargetBase     *PlayerBase
	OwnerStatus, TargetStatus InterfaceStatus
	Speed                     float64
}

// LinkExists returns true if a link between two bases has already been created
// if the link exists, it's also returned
func LinkExists(bases []*PlayerBase, a *PlayerBase, b *PlayerBase) (bool, NetworkLink) {
	var nillink NetworkLink
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
