package main

import (
	f "fmt"
	/*"net"
	"bufio"*/
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("server [numPlayers]")
	}

	numPlayers, err := strconv.Atoi(args[0])

	if err != nil {
		panic("Incorrect player number")
	}

	worldRadius := float64(numPlayers * 200)
	/*players := make([]Player, 0, numPlayers)
	bases := make([]PlayerBase, 0, numPlayers)*/
	f.Println("World size: ", worldRadius)
	locs := GeneratePlayerLocations(numPlayers, worldRadius)

	base := NewPlayerBase(locs[0], "admin", "admin")
	player := NewPlayer("roger", locs[0], base)
	err = base.AssignPlayer(player)
	f.Println(err)
	err = base.AssignPlayer(player)
	f.Println(err)
}
