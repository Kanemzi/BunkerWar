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
	players := make([]*Player, numPlayers, numPlayers)
	bases := make([]*PlayerBase, numPlayers)
	f.Println("World size: ", worldRadius)
	locs := GeneratePlayerLocations(numPlayers, worldRadius)

	for i, v := range locs {
		bases[i] = NewPlayerBase(v, "admin"+strconv.Itoa(i), "admin"+strconv.Itoa(i))
		players[i] = NewPlayer("player"+strconv.Itoa(i), v, bases[i])
		err = bases[i].AssignPlayer(players[i])
	}

	f.Println(bases)
	f.Println(players)

	f.Println(*bases[1])
	f.Println(*players[1])

}
